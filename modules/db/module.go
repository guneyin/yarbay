package db

import (
	"gorm.io/gorm"
)

const ModuleName = "db-server"

type DB struct {
	*gorm.DB
	err error
}

func (d *DB) Name() string {
	return ModuleName
}

func (d *DB) Start() error {
	if d == nil {
		return nil
	}

	return d.err
}

func (d *DB) Stop() error {
	if d == nil {
		return nil
	}

	db, err := d.DB.DB()
	if err != nil {
		return err
	}

	return db.Close()
}
