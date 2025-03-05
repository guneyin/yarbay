package nats

import "github.com/nats-io/nats.go"

func (m *NATS) Conn() *nats.Conn {
	return m.nc
}
