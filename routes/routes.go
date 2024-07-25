package routes

import (
	"event-system/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	protected := server.Group("/")
	protected.Use(middlewares.AuthMiddleware)
	protected.POST("/events", createEvent)
	protected.PUT("/events/:id", updateEvent)
	protected.DELETE("/events/:id", deleteEvent)
	protected.POST("/events/:id/register", registerForEvent)
	protected.DELETE("/events/:id/register", unregisterForEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)

}
