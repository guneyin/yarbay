package grpc

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

const (
	ModuleName = "rpc"
)

type RPC struct {
	*grpc.Server
	config Config
}

func New(config Config) *RPC {
	return &RPC{newGRPCServer(config), config}
}

func (s *RPC) Name() string {
	return ModuleName
}

func (s *RPC) Start(_ context.Context) error {
	if s == nil {
		return nil
	}

	lis, err := net.Listen("tcp", ":"+s.config.Port)
	if err != nil {
		return err
	}

	return s.Serve(lis)
}

func (s *RPC) Stop() error {
	s.Server.GracefulStop()
	return nil
}
