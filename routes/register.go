package routes

import (
	"app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find event"})
		return
	}
	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not register event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event registered"})
}
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
	}
	var event models.Event
	event.ID = eventId
	err = event.CancelRegister(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not cancel event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event cancelled"})
}
