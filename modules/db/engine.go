package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

func NewSQLiteDB(dsn string) *DB {
	return &DB{
		dialect: sqlite.Open(dsn),
	}
}

func NewPostgresDB(dsn string) *DB {
	return &DB{
		dialect: postgres.Open(dsn),
	}
}

func NewMemoryDB() *DB {
	dsn := "file::memory:?cache=shared"
	return NewSQLiteDB(dsn)
}
