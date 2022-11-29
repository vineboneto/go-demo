package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbContextPG *gorm.DB

func Connection() {

	args := []any{os.Getenv("PG_HOST"), os.Getenv("PG_USER"), os.Getenv("PG_DB"), os.Getenv("PG_PASS")}

	dns := fmt.Sprintf(`host=%s user=%s dbname=%s password=%s sslmode=disable`, args...)
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
