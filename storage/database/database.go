package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() *gorm.DB {
	conn, err := gorm.Open(sqlite.Open("storage.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	Migrate(conn)

	DB = conn

	return conn
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Host{})
}
