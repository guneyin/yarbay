package http

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

const ModuleName = "http-server"

type HTTP struct {
	*fiber.App
	config Config
}

func New(appName string, config Config) *HTTP {
	return &HTTP{
		newFiberApp(appName, config),
		config,
	}
}

func (a *HTTP) Name() string {
	return ModuleName
}

func (a *HTTP) Start(_ context.Context) error {
	if a == nil {
		return nil
	}

	return a.Listen(fmt.Sprintf(":%s", a.config.Port))
}

func (a *HTTP) Stop() error {
	if a == nil {
		return nil
	}

	return a.Shutdown()
}
