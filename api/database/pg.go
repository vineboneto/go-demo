package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dns := "host=db_go user=postgres dbname=go password=1234 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db
}
