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
