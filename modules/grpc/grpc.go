package grpc

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

const defaultTimeout = 10 * time.Second

type ServiceRegistry struct {
	Desc    grpc.ServiceDesc
	Service any
}

func newGRPCServer(config Config) *grpc.Server {
	timeout := defaultTimeout
	if config.Timeout > 0 {
		timeout = config.Timeout
	}

	svc := grpc.NewServer(
		grpc.ConnectionTimeout(timeout),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	return svc
}

func NewConnection(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
