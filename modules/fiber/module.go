package fiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

const ModuleName = "http-server"

type Fiber struct {
	*fiber.App
	port string
}

func New(config *Config) *Fiber {
	return &Fiber{
		App:  newFiberApp(config),
		port: config.Port,
	}
}

func (h *Fiber) Name() string {
	return ModuleName
}

func (h *Fiber) Start() error {
	if h == nil {
		return nil
	}

	return h.Listen(fmt.Sprintf(":%s", h.port))
}

func (h *Fiber) Stop() error {
	if h == nil {
		return nil
	}

	return h.Shutdown()
}
