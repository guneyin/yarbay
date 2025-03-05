package nats

import (
	"github.com/nats-io/nats.go"
)

const ModuleName = "nats"

type NATS struct {
	nc  *nats.Conn
	err error
}

func New(url string, options ...nats.Option) *NATS {
	nc, err := nats.Connect(url, options...)
	return &NATS{nc, err}
}

func (n *NATS) Name() string {
	return ModuleName
}

func (n *NATS) Start() error {
	if n == nil {
		return nil
	}

	return n.err
}

func (n *NATS) Stop() error {
	if n == nil {
		return nil
	}

	n.nc.Close()
	return nil
}
