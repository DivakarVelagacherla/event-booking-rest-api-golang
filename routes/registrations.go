package routes

import (
	"event-booking-rest-api-golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "event is not available"})
		return
	}

	userId := context.GetInt64("userId")

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "event is not available"})
		return
	}

	var register models.Registration
	register.EventId = eventId
	register.UserId = userId

	err = register.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to Register for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered for event"})

}

func CancelRegistration(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "event is not available"})
		return
	}

	userId := context.GetInt64("userId")

	var register models.Registration
	register.EventId = eventId
	register.UserId = userId

	err = register.Cancel()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration Cancelled"})

}
