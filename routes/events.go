package routes

import (
	"event-system/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "try later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "try later.", "error": err.Error()})
		return
	}
	event, err := models.Filter(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "try later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "try later.", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "try later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "event": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "try later.", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.Filter(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "try later.", "error": err.Error()})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "you are not allowed to update this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "try later.", "error": err.Error()})
		return
	}
	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "try later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "try later.", "error": err.Error()})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.Filter(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "try later.", "error": err.Error()})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "you are not allowed to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "try later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}
