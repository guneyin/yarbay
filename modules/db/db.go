package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormConfig = &gorm.Config{Logger: logger.Default.LogMode(logger.Error)}

func newDB(dialect gorm.Dialector) *DB {
	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		return dbErr(err)
	}

	return &DB{DB: db}
}

func dbErr(err error) *DB {
	return &DB{err: err}
}

func NewSQLiteDB(dsn string) *DB {
	return newDB(sqlite.Open(dsn))
}

func NewPostgresDB(dsn string) *DB {
	return newDB(postgres.Open(dsn))
}

func NewMemoryDB() *DB {
	dsn := "file::memory:?cache=shared"
	return NewSQLiteDB(dsn)
}

func (d *DB) WithMigrate(tables ...any) *DB {
	d.err = d.DB.AutoMigrate(tables...)
	return d
}
