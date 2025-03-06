package main

import (
	"github.com/guneyin/yarbay"
	"github.com/guneyin/yarbay/modules/db"
	"github.com/guneyin/yarbay/modules/elastic"
	"github.com/guneyin/yarbay/modules/fiber"
	"github.com/guneyin/yarbay/modules/grpc"
	natsmodule "github.com/guneyin/yarbay/modules/nats"
	"github.com/guneyin/yarbay/modules/otel"
	"github.com/guneyin/yarbay/modules/store"
	"github.com/nats-io/nats.go"
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
			Swagger: true,
		})).
		WithDB(db.NewMemoryDB()).
		WithStore(store.New()).
		WithGRPC(grpc.New(&grpc.Config{
			Port:    "8001",
			Timeout: time.Second * 10,
		})).
		WithOtel(otel.New(&otel.Config{
			AppName:   appName,
			ExportURL: "jaeger:5468",
		})).
		WithElastic(elastic.New("http://127.0.0.1:9200")).
		WithNATS(natsmodule.New(nats.DefaultURL))

	log.Fatal(app.Start())
}
