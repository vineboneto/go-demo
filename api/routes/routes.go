package routes

import (
	"vineapi/controller"
	"vineapi/middleware"

	"github.com/gin-gonic/gin"
)

func ApplyUserRoutes(router *gin.Engine) {
	router.POST("account/login", controller.AuthUser)
	router.GET("/account/me", middleware.Auth(DEFAULT), controller.LoadSelfUser)
	router.GET("/account", middleware.Auth(FLUXO), controller.LoadAllUser)
	router.POST("/account", middleware.Auth(FLUXO), controller.CreateUser)
	router.DELETE("/account", middleware.Auth(FLUXO), controller.DeleteUser)
}
