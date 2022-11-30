package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"vineapi/repo"
	"vineapi/utils"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

func Auth(role string) gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.Request.Header.Get("Authorization")

		token := strings.TrimPrefix(auth, "Bearer ")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			return
		}

		sub, err := utils.SignJWT(token, os.Getenv("JWT_SECRET"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}

		userId, _ := strconv.Atoi(sub.Subject)

		c.Set("sub", userId)

		user := repo.FindByID(userId)

		if role != "" {
			var grupos []string

			gruposJSON, _ := user.Grupos.MarshalJSON()

			json.Unmarshal(gruposJSON, &grupos)

			exist := funk.Find(grupos, func(x string) bool {
				return x == role
			})

			if exist == nil {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
				return
			}
		}

		c.Next()
	}
}
