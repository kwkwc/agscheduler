package services

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/rpc"

	"github.com/kwkwc/agscheduler"
)

type CRPCService struct {
	scheduler *agscheduler.Scheduler
	cn        *agscheduler.ClusterNode
}

func (crs *CRPCService) Register(args *agscheduler.Node, reply *agscheduler.Node) error {
	crs.cn.RPCRegister(args, reply)
	return nil
}

func (crs *CRPCService) Ping(args *agscheduler.Node, reply *agscheduler.Node) error {
	crs.cn.RPCPing(args, reply)
	return nil
}

type clusterRPCService struct {
	Scheduler *agscheduler.Scheduler
	Cn        *agscheduler.ClusterNode
}

func (s *clusterRPCService) Start() error {
	crs := &CRPCService{scheduler: s.Scheduler, cn: s.Cn}
	rpc.Register(crs)
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", s.Cn.Endpoint)
	if err != nil {
		return fmt.Errorf("cluster RPC Service listen failure: %s", err)
	}

	go http.Serve(lis, nil)
	slog.Info(fmt.Sprintf("Cluster RPC Service listening at: %s", lis.Addr()))

	return nil
}