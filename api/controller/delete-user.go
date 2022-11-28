package controller

import (
	"net/http"
	"vineapi/core"
	"vineapi/repo"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	user := core.DeleteUserInput{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exist := repo.FindByID(user.UsuarioId)

	if exist.UsuarioId != 0 {
		deletedId := repo.DeleteUser(&user)

		c.JSON(http.StatusOK, gin.H{"id": deletedId})
		return
	}

	c.JSON(http.StatusNoContent, "")
}
