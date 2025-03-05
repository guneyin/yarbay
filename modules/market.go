package modules

import (
	"github.com/guneyin/yarbay/modules/db"
	"github.com/guneyin/yarbay/modules/elastic"
	"github.com/guneyin/yarbay/modules/fiber"
	"github.com/guneyin/yarbay/modules/grpc"
	"github.com/guneyin/yarbay/modules/nats"
	"github.com/guneyin/yarbay/modules/otel"
	"github.com/guneyin/yarbay/modules/store"
)

type Market interface {
	HTTP() *fiber.Fiber
	DB() *db.DB
	Store() *store.Store
	NATS() *nats.NATS
	RPC() *grpc.GRPC
	Otel() *otel.Otel
	Elastic() *elastic.Elastic
}

func (mc *Controller) HTTP() *fiber.Fiber {
	return mc.GetModule(fiber.ModuleName).(*fiber.Fiber)
}

func (mc *Controller) DB() *db.DB {
	return mc.GetModule(db.ModuleName).(*db.DB)
}

func (mc *Controller) Store() *store.Store {
	return mc.GetModule(store.ModuleName).(*store.Store)
}

func (mc *Controller) NATS() *nats.NATS {
	return mc.GetModule(nats.ModuleName).(*nats.NATS)
}

func (mc *Controller) RPC() *grpc.GRPC {
	return mc.GetModule(grpc.ModuleName).(*grpc.GRPC)
}

func (mc *Controller) Otel() *otel.Otel {
	return mc.GetModule(otel.ModuleName).(*otel.Otel)
}

func (mc *Controller) Elastic() *elastic.Elastic {
	return mc.GetModule(elastic.ModuleName).(*elastic.Elastic)
}
