# yarbay

Microservice Application SDK

## W.I.P
This project is heavily in development, please Do NOT use in production.

## Components

- `Database [SQLite, PostgreSQL]`
- `HTTP Router [Fiber]`
- `Vector Database [Elasticsearch]`
- `Tracing [OpenTelemetry]`
- `Message Broker [NATS]`
- `GRPC Server`
- `In-Memory Store`

## Installation

```bash
  go get github.com/guneyin/yarbay
```

## Usage/Examples

```go
package main

import (
	"fmt"
	"github.com/guneyin/yarbay"
	"github.com/guneyin/yarbay/modules/db"
	"github.com/guneyin/yarbay/modules/elastic"
	"github.com/guneyin/yarbay/modules/fiber"
	"github.com/guneyin/yarbay/modules/grpc"
	"github.com/guneyin/yarbay/modules/nats"
	"github.com/guneyin/yarbay/modules/otel"
	"github.com/guneyin/yarbay/modules/store"
	"log"
	"time"
)

const (
	appName  = "yarbay-demo-app"
	appTitle = "Yarbay Demo App"
)

func main() {
	app := yarbay.NewApp(&yarbay.Config{
		Name:  appName,
		Title: appTitle,
	}).
		WithFiber(fiber.New(&fiber.Config{
			AppName: appName,
			Port:    "8000",
			Timeout: time.Second * 30,
		}).WithSwagger(&fiber.SwaggerConfig{
			HostURL:  "127.0.0.1:8000",
			BasePath: "/",
			FilePath: "./docs/swagger.json",
			Path:     "/docs",
			Title:    fmt.Sprintf("%s - API Documentation", appTitle),
		})).
		WithDB(db.NewMemoryDB().WithMigrate(tables...)).
		WithStore(store.New()).
		WithGRPC(grpc.New(&grpc.Config{
			Port:    "8001",
			Timeout: time.Second * 10,
		})).
		WithOtel(otel.New(&otel.Config{
			AppName:   appName,
			ExportURL: "127.0.0.1:5468",
		})).
		WithElastic(elastic.New("http://127.0.0.1:9200")).
		WithNATS(nats.New("nats://127.0.0.1:4222"))

	service := NewService()
	app.RegisterService(service)

	log.Fatal(app.Start())
}
```

## Test Application

TestApp is wrapper of App with default modules and helper functions. It's useful to build integration tests via [Testcontainers](https://testcontainers.com)

### Usage

#### SetupTest

```go
package main

import (
	"github.com/guneyin/yarbay"
	"github.com/guneyin/yarbay/modules/db"
	"github.com/guneyin/yarbay/modules/fiber"
	"github.com/guneyin/yarbay/modules/nats"
)

var (
	testApp     *yarbay.TestApp
	testService *Service
)

func setupTest() error {
	testApp = yarbay.NewTestApp().
		WithFiber(fiber.NewTest()).
		WithDB(db.NewMemoryDB().WithMigrate(tables...)).
		WithNATS(nats.NewTest())

	testService = NewService()
	testApp.RegisterService(testService)
	
	return nil
}
```
#### IntegrationTest

```go
func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	tapp := NewTestApp()
	err := tapp.RunTests(t,
		test1,
		test2,
	)
	assert.NoError(t, err)
}
```

