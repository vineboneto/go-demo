package main

import (
	"log"
	"vineapi/database"
	"vineapi/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	router.Use(gin.Recovery())

	database.Connection()

	routes.ApplyUserRoutes(router)

	log.Println("Started at port 3333")
	log.Fatal(router.Run(":3333"))
}
