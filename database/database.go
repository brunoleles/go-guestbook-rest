package database

import (
	"log"

	"main/support"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _db *gorm.DB

func InitDB() (*gorm.DB, error) {
	if _db != nil {
		return _db, nil
	}

	dbFile, _ := support.BasePath("data.sqllite")
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error openning initializing DB: \"%s\"", err.Error())
	}

	_db = db

	return _db, nil
}

func DB() *gorm.DB {
	if _db == nil {
		log.Fatal("Database is not initialized")
	}
	return _db
}
