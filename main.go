package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.GET("/", welcomePage)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Get Events Page"})
}

func welcomePage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Home Page"})
}
