package main

import (
	"fmt"
	"log"
	"time"
	"vineapi/database"
	"vineapi/repo"
	"vineapi/utils"

	"github.com/gofiber/fiber/v2"
)

type Authentication struct {
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha"`
}

func main() {

	defer utils.TimeExec(time.Now())

	app := fiber.New()

	database.Connection()

	app.Post("account/login", func(c *fiber.Ctx) error {
		user := new(Authentication)

		c.BodyParser(user)

		fmt.Println(user)

		errors := utils.Validator(user)

		if errors != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(errors)
		}

		output := repo.FindEmail(user.Email)

		if output.Id != 0 {
			validatePass := utils.CompareHash(output.Senha, user.Senha)

			if validatePass {
				// Gerar jwt

				return c.Status(200).JSON(fiber.Map{"id": output.Id})
			}
		}

		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{"message": "Unauthorized"})
	})

	app.Get("/account", func(c *fiber.Ctx) error {

		query := &repo.LoadUsersInput{}

		c.QueryParser(query)

		output := repo.LoadUser(query)

		return c.JSON(output)
	})

	app.Post("/account", func(c *fiber.Ctx) error {

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

	app.Delete("/account", func(c *fiber.Ctx) error {
		user := &repo.DeleteUserInput{}

		c.BodyParser(&user)

		deletedId := repo.DeleteUser(user)

		return c.JSON(fiber.Map{"id": deletedId})
	})

	log.Fatal(app.Listen(":3333"))
}
