package main

import (
	"log"
	"time"
	"vineapi/database"
	"vineapi/repo"
	"vineapi/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	defer utils.TimeExec(time.Now())

	app := fiber.New()

	database.Connection()

	app.Get("/users", func(c *fiber.Ctx) error {

		query := &repo.LoadUsersInput{}

		c.QueryParser(query)

		output := repo.LoadUser(query)

		return c.JSON(output)
	})

	app.Post("/users", func(c *fiber.Ctx) error {

		user := &repo.CreateUserInput{}

		c.BodyParser(user)

		errors := utils.Validator(user)

		if errors != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(errors)
		}

		output := repo.FindEmail(user.Email)

		if output.Id != 0 {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": "Email já cadastrado"})
		}

		hasher := utils.GenerateHash(user.Senha)

		user.Senha = hasher

		newId := repo.CreateUser(user)
		return c.JSON(fiber.Map{"id": newId})
	})

	log.Fatal(app.Listen(":3333"))
}
