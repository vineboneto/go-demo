package controller

import (
	"net/http"
	"vineapi/repo"

	"github.com/gin-gonic/gin"
)

func LoadAllUser(c *gin.Context) {

	query := &repo.LoadUsersInput{}

	c.BindQuery(query)

	output := repo.LoadUser(query)

	c.JSON(http.StatusOK, output)
}
