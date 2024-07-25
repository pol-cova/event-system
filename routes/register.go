package routes

import (
	"event-system/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "try later.", "error": err.Error()})
		return
	}

	event, err := models.Filter(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "try later.", "error": err.Error()})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "try later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "registered successfully"})
}

func unregisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "try later.", "error": err.Error()})
		return
	}
	var event models.Event
	event.ID = eventId
	err = event.Unregister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "try later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "unregistered successfully"})
}
