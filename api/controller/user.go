package controller

import (
	"net/http"
	"os"
	"strconv"
	"time"
	"vineapi/repo"
	"vineapi/utils"

	"github.com/gin-gonic/gin"
)

type Authentication struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

func AuthUser(c *gin.Context) {
	var authRequest Authentication

	if err := c.ShouldBindJSON(&authRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output := repo.FindByEmail(authRequest.Email)

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
}

func CreateUser(c *gin.Context) {

	user := &repo.CreateUserInput{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output := repo.FindByEmail(user.Email)

	if output.Id != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email já cadastrado"})
		return
	}

	hasher := utils.GenerateHash(user.Senha)

	user.Senha = hasher

	newId := repo.CreateUser(user)
	c.JSON(http.StatusCreated, gin.H{"id": newId})
}

func DeleteUser(c *gin.Context) {
	user := repo.DeleteUserInput{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exist := repo.FindByID(user.Id)

	if exist.Id != 0 {
		deletedId := repo.DeleteUser(&user)

		c.JSON(http.StatusOK, gin.H{"id": deletedId})
		return
	}

	c.JSON(http.StatusNoContent, "")
}

func LoadAllUser(c *gin.Context) {

	query := repo.LoadUsersInput{}

	c.BindQuery(&query)

	output := repo.LoadUser(&query)

	c.JSON(http.StatusOK, output)
}
