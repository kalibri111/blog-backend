package connection

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dsn = "ebadatabase.db"

var DatabaseConnectionGlobal *gorm.DB

func NewSQLiteConnection() error {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		//DisableForeignKeyConstraintWhenMigrating: true,
	})

	//errMigrate := db.AutoMigrate(&model.User{}, &model.Article{}, &model.Photo{})
	//if errMigrate != nil {
	//	return errMigrate
	//}

	//if db.Migrator().CreateTable(model.User{}) != nil {
	//	log.Fatal("creation User failed")
	//}
	//if db.Migrator().CreateTable(model.Article{}) != nil {
	//	log.Fatal("creation Article failed")
	//}
	//if db.Migrator().CreateTable(model.Photo{}) != nil {
	//	log.Fatal("creation Photo failed")
	//}

	DatabaseConnectionGlobal = db

	return err
}
