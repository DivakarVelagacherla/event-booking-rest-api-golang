package routes

import (
	"event-booking-rest-api-golang/models"
	"event-booking-rest-api-golang/utils"
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

func login(context *gin.Context) {
	var user models.User

	context.ShouldBindJSON(&user)

	err := user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Bad Credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt generate token"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Login Successfull", "token": token})

}
