package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/kwkwc/agscheduler"
)

func printMsg(j agscheduler.Job) {
	slog.Info(fmt.Sprintf("Run %s %s\n", j.Name, j.Args))
}

func runExample(s *agscheduler.Scheduler) {
	agscheduler.RegisterFuncs(printMsg)

	job1 := agscheduler.Job{
		Name:     "Job1",
		Type:     agscheduler.TYPE_INTERVAL,
		Interval: "2s",
		Timezone: "Asia/Shanghai",
		Func:     printMsg,
		Args:     map[string]any{"arg1": "1", "arg2": "2", "arg3": "3"},
	}
	job1, _ = s.AddJob(job1)
	slog.Info(fmt.Sprintf("Scheduler add %s %s.\n\n", job1.Name, job1))

	job2 := agscheduler.Job{
		Name:     "Job2",
		Type:     agscheduler.TYPE_CRON,
		CronExpr: "*/1 * * * *",
		FuncName: "main.printMsg",
		Args:     map[string]any{"arg4": "4", "arg5": "5", "arg6": "6", "arg7": "7"},
	}
	job2, _ = s.AddJob(job2)
	slog.Info(fmt.Sprintf("Scheduler add %s %s.\n\n", job2.Name, job2))

	s.Start()
	slog.Info("Scheduler start.\n\n")

	job3 := agscheduler.Job{
		Name:     "Job3",
		Type:     agscheduler.TYPE_DATETIME,
		StartAt:  "2023-09-22 07:30:08",
		Timezone: "America/New_York",
		Func:     printMsg,
		Args:     map[string]any{"arg8": "8", "arg9": "9"},
	}
	job3, _ = s.AddJob(job3)
	slog.Info(fmt.Sprintf("Scheduler add %s %s.\n\n", job3.Name, job3))

	slog.Info("Sleep 10s......\n\n")
	time.Sleep(10 * time.Second)

	job1, _ = s.GetJob(job1.Id)
	slog.Info(fmt.Sprintf("Scheduler get %s %s.\n\n", job1.Name, job1))

	job2.Type = agscheduler.TYPE_INTERVAL
	job2.Interval = "4s"
	job2, _ = s.UpdateJob(job2)
	slog.Info(fmt.Sprintf("Scheduler update %s %s.\n\n", job2.Name, job2))

	slog.Info("Sleep 8s......")
	time.Sleep(8 * time.Second)

	job1, _ = s.PauseJob(job1.Id)
	slog.Info(fmt.Sprintf("Scheduler pause %s.\n\n", job1.Name))

	slog.Info("Sleep 6s......\n\n")
	time.Sleep(6 * time.Second)

	job1, _ = s.ResumeJob(job1.Id)
	slog.Info(fmt.Sprintf("Scheduler resume %s.\n\n", job1.Name))

	s.DeleteJob(job2.Id)
	slog.Info(fmt.Sprintf("Scheduler delete %s.\n\n", job2.Name))

	slog.Info("Sleep 6s......\n\n")
	time.Sleep(6 * time.Second)

	s.Stop()
	slog.Info("Scheduler stop.\n\n")

	slog.Info("Sleep 3s......\n\n")
	time.Sleep(3 * time.Second)

	s.Start()
	slog.Info("Scheduler start.\n\n")

	slog.Info("Sleep 4s......\n\n")
	time.Sleep(4 * time.Second)

	s.DeleteAllJobs()
	slog.Info("Scheduler delete all jobs.\n\n")
}
