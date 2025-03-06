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

func (f *Fiber) Name() string {
	return ModuleName
}

func (f *Fiber) Start() error {
	if f == nil {
		return nil
	}

	return f.Listen(fmt.Sprintf(":%s", f.port))
}

func (f *Fiber) Stop() error {
	if f == nil {
		return nil
	}

	return f.Shutdown()
}
