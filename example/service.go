package main

import (
	"github.com/guneyin/yarbay/modules"
	"github.com/guneyin/yarbay/modules/db"
)

type Service struct {
	db *db.DB
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Bootstrap(m modules.Market) {
	s.db = m.DB()
}
