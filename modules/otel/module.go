package otel

import (
	"context"
	"go.opentelemetry.io/otel/sdk/trace"
	"time"
)

const ModuleName = "open-telemetry"

type Otel struct {
	*trace.TracerProvider
	config *Config
}

func New(config *Config) *Otel {
	return &Otel{config: config}
}

func (o *Otel) Name() string {
	return ModuleName
}

func (o *Otel) Start() error {
	if o == nil {
		return nil
	}

	tp, err := newTraceProvider(o.config.ExportURL, o.config.AppName)
	if err != nil {
		return err
	}

	o.TracerProvider = tp
	return nil
}

func (o *Otel) Stop() error {
	if o == nil {
		return nil
	}

	if o.TracerProvider == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return o.TracerProvider.Shutdown(ctx)
}
