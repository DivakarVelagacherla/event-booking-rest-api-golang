package routes

import (
	"event-booking-rest-api-golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func getEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request Error"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to Get Event, Internal Server Error"})
		return
	}

	context.JSON(http.StatusOK, event)
}
