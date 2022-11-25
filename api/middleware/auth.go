package middleware

import (
	"net/http"
	"os"
	"strings"
	"vineapi/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.Request.Header.Get("Authorization")

		token := strings.TrimPrefix(auth, "Bearer ")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			return
		}

		sub, err := utils.SignJWT(token, os.Getenv("JWT_SECRET"))

		c.Set("sub", sub)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}

		// Buscar usuário no banco com os grupos de acesso

		c.Next()
	}
}
