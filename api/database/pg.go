package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbContextPG *gorm.DB

func Connection() {
	dns := "host=localhost user=postgres dbname=go password=1234 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbl_",
			NoLowerCase:   false,
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	dbContextPG = db
}

func GetPG() *gorm.DB {
	return dbContextPG
}
