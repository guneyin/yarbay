package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const ModuleName = "db-server"

type DB struct {
	*gorm.DB
	dialect gorm.Dialector
}

func (d *DB) Name() string {
	return ModuleName
}

func (d *DB) Start() error {
	if d == nil {
		return nil
	}

	db, err := gorm.Open(d.dialect, &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		return err
	}

	d.DB = db
	return err
}

func (d *DB) Stop() error {
	if d == nil {
		return nil
	}

	sdb, err := d.DB.DB()
	if err != nil {
		return err
	}

	return sdb.Close()
}
