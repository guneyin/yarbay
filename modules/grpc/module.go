package grpc

import (
	"google.golang.org/grpc"
	"net"
)

const ModuleName = "rpc"

type GRPC struct {
	*grpc.Server
	port string
}

func New(config *Config) *GRPC {
	return &GRPC{Server: newGRPCServer(config), port: config.Port}
}

func (g *GRPC) Name() string {
	return ModuleName
}

func (g *GRPC) Start() error {
	if g == nil {
		return nil
	}

	lis, err := net.Listen("tcp", ":"+g.port)
	if err != nil {
		return err
	}

	return g.Serve(lis)
}

func (g *GRPC) Stop() error {
	if g == nil {
		return nil
	}

	g.Server.GracefulStop()
	return nil
}
