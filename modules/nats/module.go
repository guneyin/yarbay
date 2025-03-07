package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	natsContainer "github.com/testcontainers/testcontainers-go/modules/nats"
	"time"
)

const ModuleName = "nats"

type NATS struct {
	*nats.Conn
	container *natsContainer.NATSContainer
	err       error
}

func New(url string, options ...nats.Option) *NATS {

	nc, err := nats.Connect(url, options...)
	if err != nil {
		return newErr(err)
	}
	
	return &NATS{nc, nil, nil}
}

func NewTest() *NATS {
	container, err := natsContainer.Run(context.Background(), "nats:2.9")
	if err != nil {
		return newErr(err)
	}

	url, err := container.ConnectionString(context.Background())
	if err != nil {
		return newErr(err)
	}

	n := New(url)
	n.container = container
	return n
}

func newErr(err error) *NATS {
	return &NATS{nil, nil, err}
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

	n.Conn.Close()

	timeOut := 5 * time.Second
	if n.container != nil {
		return n.container.Stop(context.Background(), &timeOut)
	}

	return nil
}
