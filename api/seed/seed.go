package main

import (
	"vine-api/database"
	model "vine-api/models"

	"github.com/bxcodec/faker/v3"
)

func main() {
	database.Connection()

	database.DB.Exec(`
		CREATE TABLE IF NOT EXISTS tbl_usuario (
			id SERIAL PRIMARY KEY NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			first_name VARCHAR(255) NOT NULL		
		);
	`)

	for i := 0; i < 1000; i++ {
		database.DB.Table("tbl_usuario").Create(&model.User{FirstName: faker.FirstName(), LastName: faker.LastName()})
	}
}
