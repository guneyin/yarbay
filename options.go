package yarbay

import (
	"github.com/guneyin/yarbay/modules/db"
	"github.com/guneyin/yarbay/modules/elastic"
	"github.com/guneyin/yarbay/modules/fiber"
	"github.com/guneyin/yarbay/modules/grpc"
	"github.com/guneyin/yarbay/modules/nats"
	"github.com/guneyin/yarbay/modules/otel"
	"github.com/guneyin/yarbay/modules/store"
)

func (a *App) WithFiber(fiber *fiber.Fiber) *App {
	a.mc.RegisterModule(fiber)
	return a
}

func (a *App) WithDB(db *db.DB) *App {
	a.mc.RegisterModule(db)
	return a
}

func (a *App) WithStore(store *store.Store) *App {
	a.mc.RegisterModule(store)
	return a
}

func (a *App) WithNATS(nats *nats.NATS) *App {
	a.mc.RegisterModule(nats)
	return a
}

func (a *App) WithGRPC(grpc *grpc.GRPC) *App {
	a.mc.RegisterModule(grpc)
	return a
}

func (a *App) WithOtel(otel *otel.Otel) *App {
	a.mc.RegisterModule(otel)
	return a
}

func (a *App) WithElastic(elastic *elastic.Elastic) *App {
	a.mc.RegisterModule(elastic)
	return a
}
