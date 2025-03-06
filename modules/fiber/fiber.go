package fiber

import (
	"fmt"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/swaggo/swag"
	"strings"

	"os"
	"time"
)

const defaultTimeout = 30 * time.Second

func newFiberApp(config *Config) *fiber.App {
	timeout := defaultTimeout
	if config.Timeout > 0 {
		timeout = config.Timeout
	}

	app := fiber.New(fiber.Config{
		ServerHeader:          fmt.Sprintf("%s %s", config.AppName, ModuleName),
		BodyLimit:             16 * 1024 * 1024,
		AppName:               config.AppName,
		ReadTimeout:           timeout,
		WriteTimeout:          timeout,
		DisableStartupMessage: true,
	})

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(otelfiber.Middleware(
		otelfiber.WithServerName(config.AppName),
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
			Title:       fmt.Sprintf("%s API documentation", config.AppName),
			CacheAge:    0,
		}

		if strings.TrimSpace(config.ExternalIP) != "" {
			if swg, ok := swag.GetSwagger("swagger").(*swag.Spec); ok {
				swg.Host = config.ExternalIP
			}
		}

		_, err := os.Stat(swaggerConfig.FilePath)
		if err == nil {
			app.Use(swagger.New(swaggerConfig))
		}
	}

	return app
}
