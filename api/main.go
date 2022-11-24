package main

import (
	"log"
	"time"
	"vine-api/database"
	model "vine-api/models"
	"vine-api/repo"

	"github.com/gofiber/fiber/v2"
)

func main() {

	defer TimeExec(time.Now())

	app := fiber.New()

	database.Connection()

	app.Get("/users", func(c *fiber.Ctx) error {

		query := &repo.QueryLoadUsers{}

		c.QueryParser(query)

		output := repo.LoadUser(query)

		return c.JSON(output)
	})

	app.Post("/users", func(c *fiber.Ctx) error {

		user := &model.User{}

		c.BodyParser(user)

		errors := Validator(user)

		if errors != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(errors)
		}

		// Verificar se email já existe
		// Gerar hash de senha
		// Cadastrar no banco

		repo.CreateUser(user)
		return c.SendStatus(200)
	})

	log.Fatal(app.Listen(":3333"))
}
