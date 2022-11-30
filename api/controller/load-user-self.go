package controller

import (
	"net/http"
	"vineapi/repo"

	"github.com/gin-gonic/gin"
)

func LoadSelfUser(c *gin.Context) {

	sub, _ := c.Get("sub")

	output := repo.FindByID(sub.(int))

	c.JSON(http.StatusOK, output)
}
