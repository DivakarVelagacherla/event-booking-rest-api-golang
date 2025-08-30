package main

import (
	"event-booking-rest-api-golang/database"
	"event-booking-rest-api-golang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initializing DB
	database.Init()

	// Creating Server Engine
	server := gin.Default()

	// Register Routes
	routes.RegisterRoutes(server)

	// Running the server
	server.Run(":8080")
}
