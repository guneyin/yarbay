package yarbay

import (
	"github.com/guneyin/yarbay/modules"
	"google.golang.org/grpc"
)

type Service interface {
	Boostrap(m modules.Market)
}

func (a *App) RegisterService(s Service) {
	s.Boostrap(a.Market())
}

func (a *App) RegisterRPCService(desc grpc.ServiceDesc, service any) {
	if t, ok := service.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	a.Market().RPC().RegisterService(&desc, service)
}
