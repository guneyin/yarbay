package yarbay

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
)

type Ctx struct {
	*fiber.Ctx
}

func (c *Ctx) Context() context.Context {
	return c.Ctx.UserContext()
}

func (c *Ctx) SpanFromContext() trace.Span {
	return trace.SpanFromContext(c.Context())
}
