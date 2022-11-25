package main

import (
	"fmt"
	"vineapi/database"
	"vineapi/repo"

	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {

	database.Connection()

	users := []repo.CreateUserInput{}

	for i := 0; i < 1000; i++ {
		senhaHash, _ := bcrypt.GenerateFromPassword([]byte(faker.Password()), bcrypt.DefaultCost)

		users = append(users, repo.CreateUserInput{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			Senha:     string(senhaHash),
		})

	}

	database.DB.
		Table("tbl_usuario").
		CreateInBatches(users, len(users))

	database.DB.Transaction(func(tx *gorm.DB) error {
		senhaHash, _ := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)

		grupo := struct {
			Id   int
			Nome string
		}{Nome: "FLUXO", Id: 1}

		usuario := &repo.CreateUserInput{
			FirstName: "Vinicius",
			LastName:  "Boneto",
			Email:     "vineboneto@gmail.com",
			Senha:     string(senhaHash),
		}

		tx.
			Table("tbl_grupoacesso").
			Create(&grupo)

		tx.
			Table("tbl_usuario").
			Create(&usuario)

		fmt.Println(usuario)

		grupoUsuario := struct {
			IdGrupoacesso int
			IdUsuario     int
		}{IdGrupoacesso: 1, IdUsuario: usuario.Id}

		tx.
			Table("tbl_grupoacesso_usuario").
			Create(grupoUsuario)

		return nil
	})

}
