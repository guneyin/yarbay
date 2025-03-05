package nats

import (
	"context"
	"github.com/nats-io/nats.go"
)

const ModuleName = "nats"

type NATS struct {
	nc  *nats.Conn
	err error
}

func New(url ...string) *NATS {
	url = append(url, nats.DefaultURL)
	nc, err := nats.Connect(url[0])
	return &NATS{nc, err}
}

func (m *NATS) Name() string {
	return ModuleName
}

func (m *NATS) Start(_ context.Context) error {
	if m == nil {
		return nil
	}

	return m.err
}

func (m *NATS) Stop() error {
	if m == nil {
		return nil
	}

	m.nc.Close()
	return nil
}
