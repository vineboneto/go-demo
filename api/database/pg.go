package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbContextPG *gorm.DB

func Connection() {
	dns := "host=localhost user=postgres dbname=go password=1234 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dbContextPG = db
}

func GetPG() *gorm.DB {
	return dbContextPG
}
