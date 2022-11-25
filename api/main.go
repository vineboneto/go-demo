package main

import (
	"fmt"
	"log"
	"vineapi/controller"
	"vineapi/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	router.Use(gin.Recovery())

	database.Connection()

	router.POST("account/login", controller.AuthUser)
	router.GET("/account", controller.LoadAllUser)
	router.POST("/account", controller.CreateUser)
	router.DELETE("/account", controller.DeleteUser)

	fmt.Println("Started at port 3333")
	log.Fatal(router.Run(":3333"))
}
