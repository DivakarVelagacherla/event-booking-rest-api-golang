package routes

import (
	"event-booking-rest-api-golang/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/", welcomePage)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authentication)
	authenticated.POST("/events", createEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.PUT("/events/:id", updateEvent)

	server.POST("/signup", signup)

	server.POST("/login", login)
}
