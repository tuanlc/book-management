package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbClient *gorm.DB
)

func Connect() {
	databaseDNS := os.Getenv("DATABASE_DNS")

	if databaseDNS == "" {
		databaseDNS = "host=localhost user=admin password=admin dbname=book-management port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	}
	db, err := gorm.Open(postgres.Open(databaseDNS), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dbClient = db
}

func GetDBClient() *gorm.DB {
	return dbClient
}
