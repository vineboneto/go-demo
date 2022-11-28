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

	if output.UsuarioId != 0 {
		validatePass := utils.CompareHash(output.Senha, authRequest.Senha)

		if validatePass {

			jwt, err := utils.GenerateJWT(strconv.Itoa(output.UsuarioId), os.Getenv("JWT_SECRET"), time.Minute*15)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}

			refresh, err := utils.GenerateJWT(strconv.Itoa(output.UsuarioId), os.Getenv("REFRESH_SECRET"), time.Hour*24)

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
