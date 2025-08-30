package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Registering GET endpoint
	server.GET("/events", getEvents)
	server.GET("/", welcomePage)
	server.GET("/events/:id", getEvent)

	// Registering POST Endpoints
	server.POST("/events", createEvent)
}
