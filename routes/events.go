package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/learn-gin/models"
)

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the events, Try again later.", "error": err})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		if event.Name == "" {
			context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the events, Try again later.", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Found", "event": event})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the events, Try again later.", "error": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not parse data.", "error": err.Error()})
		return
	}
	event.UserID = context.GetInt64("userId")
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event, Try again later.", "error": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created.", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the events, Try again later.", "error": err})
		return
	}

	event, err := models.GetEventById(eventId)
	userId := context.GetInt64("userId")

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	if err != nil {
		if event.Name == "" {
			context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the events, Try again later.", "error": err})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message: ": "Could not parse data.", "error": err.Error()})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update the event, Try again later.", "error": err.Error()})
		return
	}

	event, err = models.GetEventById(eventId)
	if err != nil {
		if event.Name == "" {
			context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the events, Try again later.", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Updated", "updatedEvent": event})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the events, Try again later.", "error": err})
		return
	}

	event, err := models.GetEventById(eventId)
	userId := context.GetInt64("userId")

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	if err != nil {
		if event.Name == "" {
			context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the events, Try again later.", "error": err.Error()})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete the events, Try again later.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted"})
}
