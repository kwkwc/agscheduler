package agscheduler

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	TYPE_DATETIME = "datetime"
	TYPE_INTERVAL = "interval"
	TYPE_CRON     = "cron"

	STATUS_RUNNING = "running"
	STATUS_PAUSED  = "paused"
)

type Job struct {
	Id          string
	Name        string
	Type        string
	StartAt     string
	EndAt       string
	Interval    string
	CronExpr    string
	Timezone    string
	Func        func(Job)
	FuncName    string
	Args        map[string]any
	LastRunTime time.Time
	NextRunTime time.Time
	Status      string
}

func (j *Job) SetId() {
	j.Id = strings.Replace(uuid.New().String(), "-", "", -1)
}

func (j *Job) LastRunTimeWithTimezone() time.Time {
	timezone, _ := time.LoadLocation(j.Timezone)

	return j.LastRunTime.In(timezone)
}

func (j *Job) NextRunTimeWithTimezone() time.Time {
	timezone, _ := time.LoadLocation(j.Timezone)

	return j.NextRunTime.In(timezone)
}

func (j Job) String() string {
	return fmt.Sprintf(
		"Job{'Id':'%s', 'Name':'%s', 'Type':'%s', 'StartAt':'%s', 'EndAt':'%s', "+
			"'Interval':'%s', 'CronExpr':'%s', 'Timezone':'%s', "+
			"'FuncName':'%s', 'Args':'%s', "+
			"'LastRunTime':'%s', 'NextRunTime':'%s', 'Status':'%s'}",
		j.Id, j.Name, j.Type, j.StartAt, j.EndAt,
		j.Interval, j.CronExpr, j.Timezone,
		j.FuncName, j.Args,
		j.LastRunTimeWithTimezone(), j.NextRunTimeWithTimezone(), j.Status,
	)
}

func StateDumps(j Job) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(j)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func StateLoads(state []byte) (Job, error) {
	var j Job
	buf := bytes.NewBuffer(state)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&j)
	if err != nil {
		return Job{}, err
	}
	return j, nil
}

var funcMap = make(map[string]func(Job))

func getFuncName(f func(Job)) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func RegisterFuncs(fs ...func(Job)) {
	for _, f := range fs {
		fName := getFuncName(f)
		funcMap[fName] = f
	}
}
