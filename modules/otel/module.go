package otel

import (
	"context"
	"go.opentelemetry.io/otel/sdk/trace"
)

const ModuleName = "open-telemetry"

type Otel struct {
	*trace.TracerProvider
	ctx         context.Context
	appName     string
	exporterURL string
}

func New(appName, exporterURL string) *Otel {
	return &Otel{appName: appName, exporterURL: exporterURL}
}

func (o *Otel) Name() string {
	return ModuleName
}

func (o *Otel) Start(ctx context.Context) error {
	if o == nil {
		return nil
	}

	o.ctx = ctx
	tp, err := newTraceProvider(ctx, o.exporterURL, o.appName)
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

	return o.TracerProvider.Shutdown(o.ctx)
}
