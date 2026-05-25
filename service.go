package yarbay

import (
	"github.com/guneyin/yarbay/modules"
)

type Service interface {
	Bootstrap(m modules.Market)
}

func (a *App) RegisterService(s Service) {
	s.Bootstrap(a.Market())
}
