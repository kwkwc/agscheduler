package services

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	"github.com/kwkwc/agscheduler"
	pb "github.com/kwkwc/agscheduler/services/proto"
)

func panicInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	defer func() {
		if err := recover(); err != nil {
			slog.Error(fmt.Sprintf("gRPC Service method: `%s`, err: `%s`", info.FullMethod, err))
		}
	}()

	resp, err = handler(ctx, req)
	return resp, err
}

type GRPCService struct {
	Scheduler *agscheduler.Scheduler

	// Default: `127.0.0.1:36360`
	Address string

	srv *grpc.Server
}

func (s *GRPCService) Start() error {
	if s.Address == "" {
		s.Address = "127.0.0.1:36360"
	}

	lis, err := net.Listen("tcp", s.Address)
	if err != nil {
		return fmt.Errorf("gRPC Service listen failure: %s", err)
	}

	cp := &ClusterProxy{Scheduler: s.Scheduler}
	s.srv = grpc.NewServer(grpc.ChainUnaryInterceptor(panicInterceptor, cp.GRPCProxyInterceptor))

	sgrs := &sGRPCService{scheduler: s.Scheduler}
	pb.RegisterSchedulerServer(s.srv, sgrs)

	if s.Scheduler.IsClusterMode() {
		cgrs := &cGRPCService{cn: agscheduler.GetClusterNode(s.Scheduler)}
		pb.RegisterClusterServer(s.srv, cgrs)
	}

	slog.Info(fmt.Sprintf("gRPC Service listening at: %s", lis.Addr()))

	go func() {
		if err := s.srv.Serve(lis); err != nil {
			slog.Error(fmt.Sprintf("gRPC Service Unavailable: %s", err))
		}
	}()

	return nil
}

func (s *GRPCService) Stop() error {
	slog.Info("gRPC Service stop")

	s.srv.Stop()

	return nil
}