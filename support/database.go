package support

import (
	"log"

	"main/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var gdb *gorm.DB

func InitDB() (*gorm.DB, error) {
	if gdb != nil {
		return gdb, nil
	}

	dbFile, _ := BasePath("data.sqllite")
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	gdb = db

	AutoMigrateAll()

	return gdb, err
}

func GetDB() *gorm.DB {
	if gdb == nil {
		log.Fatal("Database is not initialized")
	}
	return gdb
}

func AutoMigrateAll() {
	err := GetDB().AutoMigrate(
		&models.GuestbookModel{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
