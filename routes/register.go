package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/learn-gin/models"
)

func registerForEvents(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userId := context.GetInt64("userId")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the event id., Try again later.", "error": err.Error()})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the event., Try again later.", "error": err.Error()})
		return
	}

	var register models.Register
	register.EventID = eventId
	register.UserID = userId
	err = register.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event., Try again later.", "error": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered for event."})
}

func cancelRegistration(context *gin.Context) {}
