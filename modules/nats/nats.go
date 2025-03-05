package nats

import "github.com/nats-io/nats.go"

func (n *NATS) Conn() *nats.Conn {
	return n.nc
}
