package main

import (
	"event-booking-rest-api-golang/database"
	"event-booking-rest-api-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initializing DB
	database.Init()

	// Creating Server Engine
	server := gin.Default()

	// Registering GET endpoint
	server.GET("/events", getEvents)
	server.GET("/", welcomePage)

	// Registering POST Endpoints
	server.POST("/events", createEvent)

	// Running the server
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error, Try Later!"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func welcomePage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Home Page"})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt parse the request"})
		return
	}

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save to DB, Internal Server Error"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
