package yarbay

import (
	"github.com/guneyin/yarbay/modules/db"
	"github.com/guneyin/yarbay/modules/elastic"
	"github.com/guneyin/yarbay/modules/grpc"
	"github.com/guneyin/yarbay/modules/http"
	"github.com/guneyin/yarbay/modules/otel"
	"github.com/guneyin/yarbay/modules/store"
)

func (a *App) WithHttp(config http.Config) *App {
	a.mc.RegisterModule(http.New(a.config.Title, config))
	return a
}

func (a *App) WithDB(db *db.DB) *App {
	a.mc.RegisterModule(db)
	return a
}

func (a *App) WithStore() *App {
	a.mc.RegisterModule(store.New())
	return a
}

func (a *App) WithRPC(config grpc.Config) *App {
	a.mc.RegisterModule(grpc.New(config))
	return a
}

func (a *App) WithOtel(exporterURL string) *App {
	a.mc.RegisterModule(otel.New(a.config.Name, exporterURL))
	return a
}

func (a *App) WithElastic(addr string) *App {
	a.mc.RegisterModule(elastic.New(addr))
	return a
}
