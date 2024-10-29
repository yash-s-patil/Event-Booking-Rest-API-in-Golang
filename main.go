package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// Behind the scene it configures a http server, that has some basic functionality to log incoming request and automatically recover if some part of program will crash
	server := gin.Default()

	routes.RegisterRoutes(server)

	// server starts listening for incoming request on port 8080
	server.Run(":8080") // localhost:8080
}
