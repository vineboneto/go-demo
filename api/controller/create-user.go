package controller

import (
	"net/http"
	"vineapi/core"
	"vineapi/repo"
	"vineapi/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	user := &core.CreateUserInput{}

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
