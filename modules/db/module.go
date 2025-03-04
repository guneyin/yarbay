package db

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const ModuleName = "db-server"

type DB struct {
	*gorm.DB
	dialect gorm.Dialector
}

func (s *DB) Name() string {
	return ModuleName
}

func (s *DB) Start(ctx context.Context) error {
	if s == nil {
		return nil
	}

	db, err := gorm.Open(s.dialect, &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		return err
	}

	s.DB = db.WithContext(ctx)
	return err
}

func (s *DB) Stop() error {
	if s == nil {
		return nil
	}

	sdb, err := s.DB.DB()
	if err != nil {
		return err
	}

	return sdb.Close()
}
