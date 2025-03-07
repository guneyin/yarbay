package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/propagation"
	"time"
)

func newMsgWithTrace(ctx context.Context, subject string, in any) (*nats.Msg, error) {
	data, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	header := make(nats.Header)

	propagator := propagation.TraceContext{}
	propagator.Inject(ctx, propagation.HeaderCarrier(header))
	return &nats.Msg{
		Subject: subject,
		Reply:   "",
		Header:  header,
		Data:    data,
		Sub:     nil,
	}, nil
}

func (n *NATS) RequestWithTrace(ctx context.Context, subject string, in any, d time.Duration) (*nats.Msg, error) {
	msg, err := newMsgWithTrace(ctx, subject, in)
	if err != nil {
		return nil, err
	}

	return n.Conn.RequestMsg(msg, d)
}

func (n *NATS) SubscribeWithTrace(subject string, callback func(ctxCallback context.Context, msg *nats.Msg)) (*nats.Subscription, error) {
	subscription, err := n.Conn.Subscribe(subject, func(msg *nats.Msg) {
		propagator := propagation.TraceContext{}
		ctx := propagator.Extract(context.Background(), propagation.HeaderCarrier(msg.Header))
		callback(ctx, msg)
	})
	if err != nil {
		return nil, fmt.Errorf("could not subscribe to topic: %w", err)
	}

	return subscription, nil
}

func (n *NATS) PublishWithTrace(ctx context.Context, subject string, in any) error {
	msg, err := newMsgWithTrace(ctx, subject, in)
	if err != nil {
		return err
	}

	return n.Conn.PublishMsg(msg)
}
