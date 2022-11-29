package main

import (
	"vineapi/core"
	"vineapi/database"

	"github.com/bxcodec/faker/v3"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	database.Connection()

	users := []core.CreateUserInput{}

	for i := 0; i < 1000; i++ {
		senhaHash, _ := bcrypt.GenerateFromPassword([]byte(faker.Password()), bcrypt.DefaultCost)

		users = append(users, core.CreateUserInput{
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			Senha:     string(senhaHash),
		})

	}

	database.GetPG().
		Table("tbl_usuario").
		CreateInBatches(users, 100)

	database.GetPG().Transaction(func(tx *gorm.DB) error {
		senhaHash, _ := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)

		grupo := struct {
			IdGrupoacesso int
			Nome          string
		}{Nome: "FLUXO", IdGrupoacesso: 1}

		usuario := &core.CreateUserInput{
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

		grupoUsuario := struct {
			IdGrupoacesso int
			IdUsuario     int
		}{IdGrupoacesso: 1, IdUsuario: usuario.UsuarioId}

		tx.
			Table("tbl_grupoacesso_usuario").
			Create(grupoUsuario)

		return nil
	})

}
