package fiber

import (
	"encoding/json"
	"fmt"
	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"os"
	"path"
	"strings"
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

	return app
}

func (f *Fiber) Addr() string {
	return fmt.Sprintf(":%s", f.port)
}

func (f *Fiber) WithSwagger(config ...*SwaggerConfig) *Fiber {
	cfg := NewSwaggerConfig(config...)
	uiPath := path.Join(cfg.BasePath, cfg.Path)

	f.App.Get(uiPath, func(c *fiber.Ctx) error {
		specContent, err := readSwaggerContent(cfg.FilePath)
		if err != nil {
			return err
		}

		if strings.TrimSpace(cfg.HostURL) != "" {
			specContent["host"] = cfg.HostURL
		}

		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecContent: specContent,
			CustomOptions: scalar.CustomOptions{
				PageTitle: cfg.Title,
			},
		})

		if err != nil {
			fmt.Printf("%v", err)
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(htmlContent)
	})

	return f
}

func readSwaggerContent(path string) (map[string]interface{}, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	data := make(map[string]interface{}, 0)
	if err = dec.Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
