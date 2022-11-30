package routes

import (
	"vineapi/controller"
	"vineapi/core"
	"vineapi/middleware"

	"github.com/gin-gonic/gin"
)

func ApplyUserRoutes(router *gin.Engine) {
	router.POST("account/login", controller.AuthUser)
	router.GET("/account/me", middleware.Auth(core.DEFAULT), controller.LoadSelfUser)
	router.GET("/account", middleware.Auth(core.FLUXO), controller.LoadAllUser)
	router.POST("/account", middleware.Auth(core.FLUXO), controller.CreateUser)
	router.DELETE("/account", middleware.Auth(core.FLUXO), controller.DeleteUser)
}
