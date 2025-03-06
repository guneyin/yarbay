package yarbay

import (
	"github.com/guneyin/yarbay/modules"
)

type Service interface {
	Boostrap(m modules.Market)
}

func (a *App) RegisterService(s Service) {
	s.Boostrap(a.Market())
}
