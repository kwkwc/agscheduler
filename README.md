# AGScheduler

[![test](https://github.com/kwkwc/agscheduler/actions/workflows/test.yml/badge.svg)](https://github.com/kwkwc/agscheduler/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/kwkwc/agscheduler/graph/badge.svg?token=CL5P4VYQTU)](https://codecov.io/gh/kwkwc/agscheduler)
[![Go Report Card](https://goreportcard.com/badge/github.com/kwkwc/agscheduler)](https://goreportcard.com/report/github.com/kwkwc/agscheduler)
[![Go Reference](https://pkg.go.dev/badge/github.com/kwkwc/agscheduler.svg)](https://pkg.go.dev/github.com/kwkwc/agscheduler)
![GitHub release (with filter)](https://img.shields.io/github/v/release/kwkwc/agscheduler)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/kwkwc/agscheduler)
[![license](https://img.shields.io/github/license/kwkwc/agscheduler)](https://github.com/kwkwc/agscheduler/blob/main/LICENSE)

> Advanced Golang Scheduler (AGScheduler) is a task scheduling library for Golang that supports multiple scheduling types, dynamically changing and persistent jobs, remote call, and cluster

English | [简体中文](README.zh-CN.md)

## Features

- Supports three scheduling types
  - [x] One-off execution
  - [x] Interval execution
  - [x] Cron-style scheduling
- Supports multiple job store methods
  - [x] Memory
  - [x] [GROM](https://gorm.io/)(any RDBMS supported by GROM works)
  - [x] [Redis](https://redis.io/)
  - [x] [MongoDB](https://www.mongodb.com/)
  - [x] [etcd](https://etcd.io/)
- Supports remote call
  - [x] gRPC
  - [x] HTTP APIs
- Supports cluster
  - [x] Remote worker nodes

## Framework

![Framework](assets/framework.png)

## Usage

```golang
package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/kwkwc/agscheduler"
	"github.com/kwkwc/agscheduler/stores"
)

func printMsg(ctx context.Context, j agscheduler.Job) {
	slog.Info(fmt.Sprintf("Run job `%s` %s\n\n", j.FullName(), j.Args))
}

func main() {
	agscheduler.RegisterFuncs(printMsg)

	store := &stores.MemoryStore{}
	scheduler := &agscheduler.Scheduler{}
	scheduler.SetStore(store)

	job1 := agscheduler.Job{
		Name:     "Job1",
		Type:     agscheduler.TYPE_INTERVAL,
		Interval: "2s",
		Timezone: "UTC",
		Func:     printMsg,
		Args:     map[string]any{"arg1": "1", "arg2": "2", "arg3": "3"},
	}
	job1, _ = scheduler.AddJob(job1)
	slog.Info(fmt.Sprintf("%s.\n\n", job1))

	job2 := agscheduler.Job{
		Name:     "Job2",
		Type:     agscheduler.TYPE_CRON,
		CronExpr: "*/1 * * * *",
		Timezone: "Asia/Shanghai",
		FuncName: "main.printMsg",
		Args:     map[string]any{"arg4": "4", "arg5": "5", "arg6": "6", "arg7": "7"},
	}
	job2, _ = s.AddJob(job2)
	slog.Info(fmt.Sprintf("%s.\n\n", job2))

	job3 := agscheduler.Job{
		Name:     "Job3",
		Type:     agscheduler.TYPE_DATETIME,
		StartAt:  "2023-09-22 07:30:08",
		Timezone: "America/New_York",
		Func:     printMsg,
		Args:     map[string]any{"arg8": "8", "arg9": "9"},
	}
	job3, _ = s.AddJob(job3)
	slog.Info(fmt.Sprintf("%s.\n\n", job3))

	jobs, _ := s.GetAllJobs()
	slog.Info(fmt.Sprintf("Scheduler get all jobs %s.\n\n", jobs))

	scheduler.Start()

	select {}
}
```

## Register Funcs

> **_Since golang can't serialize functions, you need to register them with `RegisterFuncs` before `scheduler.Start()`_**

## gRPC

```golang
// Server
srservice := services.SchedulerRPCService{
	Scheduler: scheduler,
	Address:   "127.0.0.1:36360",
}
srservice.Start()

// Client
conn, _ := grpc.Dial("127.0.0.1:36360", grpc.WithTransportCredentials(insecure.NewCredentials()))
client := pb.NewSchedulerClient(conn)
client.AddJob(ctx, job)
```

## HTTP APIs

```golang
// Server
shservice := services.SchedulerHTTPService{
	Scheduler: scheduler,
	Address:   "127.0.0.1:36370",
}
shservice.Start()

// Client
mJob := map[string]any{...}
bJob, _ := json.Marshal(bJob)
resp, _ := http.Post("http://127.0.0.1:36370/scheduler/job", "application/json", bytes.NewReader(bJob))
```

## Cluster

```golang
// Main Node
cnMain := &agscheduler.ClusterNodeClusterNode{
	Endpoint:          "127.0.0.1:36380",
	EndpointHTTP:      "127.0.0.1:36390",
	SchedulerEndpoint: "127.0.0.1:36360",
	Queue:             "default",
}
schedulerMain.SetClusterNode(ctx, cnMain)
cserviceMain := &services.ClusterService{Cn: cnMain}
cserviceMain.Start()

// Node
cn := &agscheduler.ClusterNode{
	MainEndpoint:      "127.0.0.1:36380",
	Endpoint:          "127.0.0.1:36381",
	EndpointHTTP:      "127.0.0.1:36391",
	SchedulerEndpoint: "127.0.0.1:36361",
	Queue:             "node",
}
scheduler.SetClusterNode(ctx, cn)
cservice := &services.ClusterService{Cn: cn}
cservice.Start()
```

## Scheduler API

| gRPC Function | HTTP Method | HTTP Endpoint             |
|---------------|-------------|---------------------------|
| AddJob        | POST        | /scheduler/job            |
| GetJob        | GET         | /scheduler/job/:id        |
| GetAllJobs    | GET         | /scheduler/jobs           |
| UpdateJob     | PUT         | /scheduler/job            |
| DeleteJob     | DELETE      | /scheduler/job/:id        |
| DeleteAllJobs | DELETE      | /scheduler/jobs           |
| PauseJob      | POST        | /scheduler/job/:id/pause  |
| ResumeJob     | POST        | /scheduler/job/:id/resume |
| RunJob        | POST        | /scheduler/job/run        |
| Start         | POST        | /scheduler/start          |
| Stop          | POST        | /scheduler/stop           |

## Cluster API

| RPC Function | HTTP Method | HTTP Endpoint             |
|---------------|-------------|---------------------------|
| Nodes         | GET         | /cluster/nodes            |

## Examples

[Complete example][examples]

## Thanks

[APScheduler](https://github.com/agronholm/apscheduler/tree/3.x)

[examples]: https://github.com/kwkwc/agscheduler/tree/main/examples
