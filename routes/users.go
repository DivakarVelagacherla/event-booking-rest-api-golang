package routes

import (
	"event-booking-rest-api-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt parse the request"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save user to DB, Internal Server Error"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
