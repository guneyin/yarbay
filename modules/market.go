package modules

import (
	"github.com/guneyin/yarbay/modules/db"
	"github.com/guneyin/yarbay/modules/elastic"
	"github.com/guneyin/yarbay/modules/grpc"
	"github.com/guneyin/yarbay/modules/http"
	"github.com/guneyin/yarbay/modules/otel"
	"github.com/guneyin/yarbay/modules/store"
)

type Market interface {
	HTTP() *http.HTTP
	DB() *db.DB
	Store() *store.Store
	RPC() *grpc.RPC
	Otel() *otel.Otel
	Elastic() *elastic.Elastic
}

func (mc *Controller) HTTP() *http.HTTP {
	return mc.GetModule(http.ModuleName).(*http.HTTP)
}

func (mc *Controller) DB() *db.DB {
	return mc.GetModule(db.ModuleName).(*db.DB)
}

func (mc *Controller) Store() *store.Store {
	return mc.GetModule(store.ModuleName).(*store.Store)
}

func (mc *Controller) RPC() *grpc.RPC {
	return mc.GetModule(grpc.ModuleName).(*grpc.RPC)
}

func (mc *Controller) Otel() *otel.Otel {
	return mc.GetModule(otel.ModuleName).(*otel.Otel)
}

func (mc *Controller) Elastic() *elastic.Elastic {
	return mc.GetModule(elastic.ModuleName).(*elastic.Elastic)
}
