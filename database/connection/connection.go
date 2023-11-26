package connection

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dsn = "127.0.0.1"

var DatabaseConnectionGlobal *gorm.DB

func newSQLiteConnection() error {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	DatabaseConnectionGlobal = db

	return err
}
