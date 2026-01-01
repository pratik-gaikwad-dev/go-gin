package main

import (
	"github.com/gin-gonic/gin"
	"github.com/learn-gin/db"
	"github.com/learn-gin/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
