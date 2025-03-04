package http

import (
	"fmt"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"time"
)

const defaultTimeout = 30 * time.Second

func newFiberApp(appName string, config Config) *fiber.App {
	timeout := defaultTimeout
	if config.Timeout > 0 {
		timeout = config.Timeout
	}

	app := fiber.New(fiber.Config{
		ServerHeader:          fmt.Sprintf("%s %s", appName, ModuleName),
		BodyLimit:             16 * 1024 * 1024,
		AppName:               appName,
		ReadTimeout:           timeout,
		WriteTimeout:          timeout,
		DisableStartupMessage: true,
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(otelfiber.Middleware(
		otelfiber.WithServerName(appName),
		otelfiber.WithSpanNameFormatter(func(c *fiber.Ctx) string {
			return c.Route().Name + ": " + c.Method() + " " + c.Path()
		}),
	))

	if config.Swagger {
		swaggerConfig := swagger.Config{
			Next:        nil,
			BasePath:    "/",
			FilePath:    "./docs/swagger.json",
			FileContent: nil,
			Path:        "docs",
			Title:       fmt.Sprintf("%s API documentation", appName),
			CacheAge:    0,
		}
		app.Use(swagger.New(swaggerConfig))
	}

	return app
}
