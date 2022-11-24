package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"vineapi/database"
	"vineapi/repo"
	"vineapi/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Authentication struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

func main() {
	godotenv.Load()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(gin.Recovery())

	database.Connection()

	router.POST("account/login", func(c *gin.Context) {
		var authRequest Authentication

		if err := c.ShouldBindJSON(&authRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		output := repo.FindEmail(authRequest.Email)

		if output.Id != 0 {
			validatePass := utils.CompareHash(output.Senha, authRequest.Senha)

			if validatePass {

				jwt, err := utils.GenerateJWT(strconv.Itoa(output.Id), os.Getenv("JWT_SECRET"), time.Minute*15)

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
					return
				}

				refresh, err := utils.GenerateJWT(strconv.Itoa(output.Id), os.Getenv("REFRESH_SECRET_KEY"), time.Hour*24)

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
					return
				}

				c.JSON(http.StatusAccepted, gin.H{"token": jwt, "refresh": refresh})
				return
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	})

	// app.Get("/account", func(c *fiber.Ctx) error {

	// 	query := &repo.LoadUsersInput{}

	// 	c.QueryParser(query)

	// 	output := repo.LoadUser(query)

	// 	return c.JSON(output)
	// })

	// app.Post("/account", func(c *fiber.Ctx) error {

	// 	user := &repo.CreateUserInput{}

	// 	c.BodyParser(user)

	// 	errors := utils.Validator(user)

	// 	if errors != nil {
	// 		return c.Status(fiber.ErrBadRequest.Code).JSON(errors)
	// 	}

	// 	output := repo.FindEmail(user.Email)

	// 	if output.Id != 0 {
	// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": "Email já cadastrado"})
	// 	}

	// 	hasher := utils.GenerateHash(user.Senha)

	// 	user.Senha = hasher

	// 	newId := repo.CreateUser(user)
	// 	return c.JSON(fiber.Map{"id": newId})
	// })

	// app.Delete("/account", func(c *fiber.Ctx) error {
	// 	user := &repo.DeleteUserInput{}

	// 	c.BodyParser(&user)

	// 	deletedId := repo.DeleteUser(user)

	// 	return c.JSON(fiber.Map{"id": deletedId})
	// })

	fmt.Println("Started at port 3333")
	log.Fatal(router.Run(":3333"))
}
