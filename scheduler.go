package agscheduler

import (
	"context"
	"fmt"
	"log/slog"
	"net/rpc"
	"reflect"
	"runtime/debug"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gorhill/cronexpr"
)

var GetStore = (*Scheduler).getStore
var GetClusterNode = (*Scheduler).getClusterNode

var mutexS sync.RWMutex

// In standalone mode, the scheduler only needs to run jobs on a regular basis.
// In cluster mode, the scheduler also needs to be responsible for allocating jobs to cluster nodes.
type Scheduler struct {
	// Job store
	store Store
	// When the time is up, the scheduler will wake up.
	timer *time.Timer
	// Input is received when `stop` is called or no job in store.
	quitChan chan struct{}
	// It should not be set manually.
	isRunning bool

	// Used in cluster mode, bind to each other and the cluster node.
	clusterNode *ClusterNode
}

func (s *Scheduler) IsRunning() bool {
	mutexS.RLock()
	defer mutexS.RUnlock()

	return s.isRunning
}

// Bind the store
func (s *Scheduler) SetStore(sto Store) error {
	s.store = sto
	if err := s.store.Init(); err != nil {
		return err
	}

	return nil
}

func (s *Scheduler) getStore() Store {
	return s.store
}

// Bind the cluster node
func (s *Scheduler) SetClusterNode(ctx context.Context, cn *ClusterNode) error {
	s.clusterNode = cn
	cn.Scheduler = s
	if err := s.clusterNode.init(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Scheduler) IsClusterMode() bool {
	return s.clusterNode != nil
}

func (s *Scheduler) getClusterNode() *ClusterNode {
	return s.clusterNode
}

// Calculate the next run time, different job type will be calculated in different ways,
// when the job is paused, will return `9999-09-09 09:09:09`.
func CalcNextRunTime(j Job) (time.Time, error) {
	timezone, err := time.LoadLocation(j.Timezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("job `%s` Timezone `%s` error: %s", j.FullName(), j.Timezone, err)
	}

	if j.Status == STATUS_PAUSED {
		nextRunTimeMax, _ := GetNextRunTimeMax()
		return time.Unix(nextRunTimeMax.Unix(), 0).UTC(), nil
	}

	var nextRunTime time.Time
	switch strings.ToLower(j.Type) {
	case TYPE_DATETIME:
		nextRunTime, err = time.ParseInLocation(time.DateTime, j.StartAt, timezone)
		if err != nil {
			return time.Time{}, fmt.Errorf("job `%s` StartAt `%s` error: %s", j.FullName(), j.Timezone, err)
		}
	case TYPE_INTERVAL:
		i, err := time.ParseDuration(j.Interval)
		if err != nil {
			return time.Time{}, fmt.Errorf("job `%s` Interval `%s` error: %s", j.FullName(), j.Interval, err)
		}
		nextRunTime = time.Now().In(timezone).Add(i)
	case TYPE_CRON:
		expr, err := cronexpr.Parse(j.CronExpr)
		if err != nil {
			return time.Time{}, fmt.Errorf("job `%s` CronExpr `%s` error: %s", j.FullName(), j.CronExpr, err)
		}
		nextRunTime = expr.Next(time.Now().In(timezone))
	default:
		return time.Time{}, fmt.Errorf("job `%s` Type `%s` unknown", j.FullName(), j.Type)
	}

	return time.Unix(nextRunTime.Unix(), 0).UTC(), nil
}

func (s *Scheduler) AddJob(j Job) (Job, error) {
	if err := j.init(); err != nil {
		return Job{}, err
	}

	slog.Info(fmt.Sprintf("Scheduler add job `%s`.\n", j.FullName()))

	if err := s.store.AddJob(j); err != nil {
		return Job{}, err
	}

	return j, nil
}

func (s *Scheduler) GetJob(id string) (Job, error) {
	return s.store.GetJob(id)
}

func (s *Scheduler) GetAllJobs() ([]Job, error) {
	return s.store.GetAllJobs()
}

func (s *Scheduler) UpdateJob(j Job) (Job, error) {
	oJ, err := s.GetJob(j.Id)
	if err != nil {
		return Job{}, err
	}

	if j.Status == "" ||
		(j.Status != STATUS_RUNNING && j.Status != STATUS_PAUSED) {
		j.Status = oJ.Status
	}

	if err := j.check(); err != nil {
		return Job{}, err
	}

	nextRunTime, err := CalcNextRunTime(j)
	if err != nil {
		return Job{}, err
	}
	j.NextRunTime = nextRunTime

	lastNextWakeupInterval := s.getNextWakeupInterval()

	if err := s.store.UpdateJob(j); err != nil {
		return Job{}, err
	}

	nextWakeupInterval := s.getNextWakeupInterval()
	if nextWakeupInterval < lastNextWakeupInterval {
		s.wakeup()
	}

	return j, nil
}

func (s *Scheduler) DeleteJob(id string) error {
	slog.Info(fmt.Sprintf("Scheduler delete jobId `%s`.\n", id))

	if _, err := s.GetJob(id); err != nil {
		return err
	}

	return s.store.DeleteJob(id)
}

func (s *Scheduler) DeleteAllJobs() error {
	slog.Info("Scheduler delete all jobs.\n")

	return s.store.DeleteAllJobs()
}

func (s *Scheduler) PauseJob(id string) (Job, error) {
	slog.Info(fmt.Sprintf("Scheduler pause jobId `%s`.\n", id))

	j, err := s.GetJob(id)
	if err != nil {
		return Job{}, err
	}

	j.Status = STATUS_PAUSED

	j, err = s.UpdateJob(j)
	if err != nil {
		return Job{}, err
	}

	return j, nil
}

func (s *Scheduler) ResumeJob(id string) (Job, error) {
	slog.Info(fmt.Sprintf("Scheduler resume jobId `%s`.\n", id))

	j, err := s.GetJob(id)
	if err != nil {
		return Job{}, err
	}

	j.Status = STATUS_RUNNING

	j, err = s.UpdateJob(j)
	if err != nil {
		return Job{}, err
	}

	return j, nil
}

// Used in standalone mode.
func (s *Scheduler) _runJob(j Job) {
	f := reflect.ValueOf(FuncMap[j.FuncName].Func)
	if f.IsNil() {
		slog.Warn(fmt.Sprintf("Job `%s` Func `%s` unregistered\n", j.FullName(), j.FuncName))
	} else {
		slog.Info(fmt.Sprintf("Job `%s` is running, next run time: `%s`\n", j.FullName(), j.NextRunTimeWithTimezone().String()))
		go func() {
			timeout, err := time.ParseDuration(j.Timeout)
			if err != nil {
				e := &JobTimeoutError{FullName: j.FullName(), Timeout: j.Timeout, Err: err}
				slog.Error(e.Error())
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			ch := make(chan error, 1)
			go func() {
				defer close(ch)
				defer func() {
					if err := recover(); err != nil {
						slog.Error(fmt.Sprintf("Job `%s` run error: %s\n", j.FullName(), err))
						slog.Debug(fmt.Sprintf("%s\n", string(debug.Stack())))
					}
				}()

				f.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(j)})
			}()

			select {
			case <-ch:
				return
			case <-ctx.Done():
				slog.Warn(fmt.Sprintf("Job `%s` run timeout\n", j.FullName()))
			}
		}()
	}
}

// Used in cluster mode.
// Call the RPC API of the other node to run the `RunJob`.
func (s *Scheduler) _runJobRemote(node *ClusterNode, j Job) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				slog.Error(fmt.Sprintf("Job `%s` _runJobRemote error: %s\n", j.FullName(), err))
				slog.Debug(fmt.Sprintf("%s\n", string(debug.Stack())))
			}
		}()

		rClient, err := rpc.DialHTTP("tcp", node.Endpoint)
		if err != nil {
			slog.Error(fmt.Sprintf("Failed to connect to cluster node: `%s`, error: %s", node.Endpoint, err))
		}
		defer rClient.Close()

		var r any
		ch := make(chan error, 1)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					slog.Error(fmt.Sprintf("Job `%s` CRPCService.RunJob error: %s\n", j.FullName(), err))
					slog.Debug(fmt.Sprintf("%s\n", string(debug.Stack())))
				}
			}()

			ch <- rClient.Call("CRPCService.RunJob", j, &r)
		}()
		select {
		case err := <-ch:
			if err != nil {
				slog.Error(fmt.Sprintf("Scheduler run job `%s` remote error %s\n", j.FullName(), err))
			}
		case <-time.After(3 * time.Second):
			slog.Error(fmt.Sprintf("Scheduler run job `%s` remote timeout %s\n", j.FullName(), err))
		}
	}()
}

func (s *Scheduler) _flushJob(j Job, now time.Time) error {
	j.LastRunTime = time.Unix(now.Unix(), 0).UTC()

	if j.Type == TYPE_DATETIME {
		if j.NextRunTime.Before(now) {
			if err := s.DeleteJob(j.Id); err != nil {
				return fmt.Errorf("delete job `%s` error: %s", j.FullName(), err)
			}
		}
	} else {
		if _, err := s.UpdateJob(j); err != nil {
			return fmt.Errorf("update job `%s` error: %s", j.FullName(), err)
		}
	}

	return nil
}

func (s *Scheduler) _scheduleJob(j Job) error {
	// In standalone mode.
	if !s.IsClusterMode() {
		s._runJob(j)
	} else {
		// In cluster mode, all nodes are equal and may pick myself.
		node, err := s.clusterNode.choiceNode(j.Queues)
		if err != nil || s.clusterNode.Endpoint == node.Endpoint {
			if len(j.Queues) == 0 || slices.Contains(j.Queues, s.clusterNode.Queue) {
				s._runJob(j)
			} else {
				return fmt.Errorf("cluster node with queue `%s` does not exist", j.Queues)
			}
		} else {
			s._runJobRemote(node, j)
		}
	}

	return nil
}

func (s *Scheduler) RunJob(j Job) error {
	slog.Info(fmt.Sprintf("Scheduler run job `%s`.\n", j.FullName()))

	s._runJob(j)

	return nil
}

// Used in cluster mode.
// Select a worker node
func (s *Scheduler) ScheduleJob(j Job) error {
	slog.Info(fmt.Sprintf("Scheduler schedule job `%s`.\n", j.FullName()))

	err := s._scheduleJob(j)
	if err != nil {
		return fmt.Errorf("scheduler schedule job `%s` error: %s", j.FullName(), err)
	}

	return nil
}

func (s *Scheduler) run() {
	for {
		select {
		case <-s.quitChan:
			slog.Info("Scheduler quit.\n")
			return
		case <-s.timer.C:
			if s.IsClusterMode() && !s.clusterNode.IsMainNode() {
				s.timer.Reset(time.Second)
				continue
			}

			now := time.Now().UTC()

			js, err := s.GetAllJobs()
			if err != nil {
				slog.Error(fmt.Sprintf("Scheduler get all jobs error: %s\n", err))
				s.timer.Reset(time.Second)
				continue
			}

			// If there are ineligible job, subsequent job do not need to be checked.
			sort.Sort(JobSlice(js))
			for _, j := range js {
				if j.NextRunTime.Before(now) {
					nextRunTime, err := CalcNextRunTime(j)
					if err != nil {
						slog.Error(fmt.Sprintf("Scheduler calc next run time error: %s\n", err))
						continue
					}
					j.NextRunTime = nextRunTime

					err = s._scheduleJob(j)
					if err != nil {
						slog.Error(fmt.Sprintf("Scheduler schedule job `%s` error: %s\n", j.FullName(), err))
					}

					err = s._flushJob(j, now)
					if err != nil {
						slog.Error(fmt.Sprintf("Scheduler %s\n", err))
						continue
					}
				} else {
					break
				}
			}

			nextWakeupInterval := s.getNextWakeupInterval()
			slog.Debug(fmt.Sprintf("Scheduler next wakeup interval %s\n", nextWakeupInterval))

			s.timer.Reset(nextWakeupInterval)
		}
	}
}

// In addition to being called manually,
// it is also called after `AddJob`.
func (s *Scheduler) Start() {
	mutexS.Lock()
	defer mutexS.Unlock()

	if s.isRunning {
		slog.Info("Scheduler is running.\n")
		return
	}

	s.timer = time.NewTimer(0)
	s.quitChan = make(chan struct{}, 3)
	s.isRunning = true

	go s.run()

	slog.Info("Scheduler start.\n")
}

// In addition to being called manually,
// there is no job in store that will also be called.
func (s *Scheduler) Stop() {
	mutexS.Lock()
	defer mutexS.Unlock()

	if !s.isRunning {
		slog.Info("Scheduler has stopped.\n")
		return
	}

	s.quitChan <- struct{}{}
	s.isRunning = false

	slog.Info("Scheduler stop.\n")
}

// Dynamically calculate the next wakeup interval, avoid frequent wakeup of the scheduler
func (s *Scheduler) getNextWakeupInterval() time.Duration {
	nextRunTimeMin, err := s.store.GetNextRunTime()
	if err != nil {
		slog.Error(fmt.Sprintf("Scheduler get next wakeup interval error: %s\n", err))
		nextRunTimeMin = time.Now().UTC().Add(1 * time.Second)
	}

	now := time.Now().UTC()
	nextWakeupInterval := nextRunTimeMin.Sub(now)
	if nextWakeupInterval < 0 {
		nextWakeupInterval = time.Second
	}

	return nextWakeupInterval
}

func (s *Scheduler) wakeup() {
	if s.timer != nil {
		s.timer.Reset(0)
	}
}

func (s *Scheduler) Info() map[string]any {
	info := map[string]any{
		"cluster_main_node": map[string]any{},
		"is_cluster_mode":   s.IsClusterMode(),
		"is_running":        s.IsRunning(),
		"version":           Version,
	}

	if s.IsClusterMode() {
		info["cluster_main_node"] = map[string]any{
			"endpoint_main": s.clusterNode.GetEndpointMain(),
			"endpoint":      s.clusterNode.Endpoint,
			"endpoint_grpc": s.clusterNode.EndpointGRPC,
			"endpoint_http": s.clusterNode.EndpointHTTP,
			"mode":          s.clusterNode.Mode,
		}
	}

	return info
}
