package main

import (
	"log"
	"vineapi/database"
	"vineapi/repo"

	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	database.Connection()

	database.DB.Exec(`
		CREATE TABLE IF NOT EXISTS tbl_usuario (
			id SERIAL PRIMARY KEY NOT NULL,
			email VARCHAR (255) NOT NULL UNIQUE,
			senha VARCHAR (255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			first_name VARCHAR(255) NOT NULL		
		);
	`)

	for i := 0; i < 1000; i++ {
		senhaHash, err := bcrypt.GenerateFromPassword([]byte(faker.Password()), bcrypt.DefaultCost)

		if err != nil {
			log.Fatal(err)
		}

		database.DB.
			Table("tbl_usuario").
			Create(&repo.CreateUserInput{
				FirstName: faker.FirstName(),
				LastName:  faker.LastName(),
				Email:     faker.Email(),
				Senha:     string(senhaHash),
			})
	}
}
