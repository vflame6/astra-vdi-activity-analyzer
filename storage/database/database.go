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

func CreateHost(hostname, secret string) error {
	host := &Host{
		Hostname: hostname,
		Secret:   secret,
	}
	return DB.Create(host).Error
}

func SelectHost(hostname string) (*Host, error) {
	var host Host
	err := DB.First(&host, "hostname = ?", hostname).Error
	if err != nil {
		return nil, err
	}
	return &host, nil
}
