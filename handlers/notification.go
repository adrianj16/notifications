package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notifications/models"
	"notifications/services"
)

type NotificationHandlers struct {
	notificationService *services.NotificationService
}

func NewNotificationHandlers(notificationService *services.NotificationService) *NotificationHandlers {
	return &NotificationHandlers{
		notificationService: notificationService,
	}
}

func (h *NotificationHandlers) GetNotifications(c *gin.Context) {
	c.JSON(http.StatusOK, h.GetNotifications)
}

func (h *NotificationHandlers) GetNotificationById(c *gin.Context) {
	notificationId := c.Param("notificationId")
	notification := h.notificationService.GetNotificationById(notificationId)
	if notification == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "notification not found"})
		return
	}
	c.JSON(http.StatusOK, notification)
}

func (h *NotificationHandlers) CreateNotification(c *gin.Context) {
	var notification models.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.notificationService. CreateNotification(notification)
	c.JSON(http.StatusOK, gin.H{"message": "notification created"})
}

func (h *NotificationHandlers) DeleteNotificationById(c *gin.Context) {
	notificationId := c.Param("notificationId")
	h.notificationService.DeleteNotificationById(notificationId)
	c.JSON(200, gin.H{"message": "notification deleted"})
}

func (h *NotificationHandlers) NotificationRestriction(c *gin.Context) {
	notificationId := c.Param("notificationId")
	if notificationId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "notification_id is empty"})
	}

	var restriction models.NotificationRestriction
	if err := c.ShouldBindJSON(&restriction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.notificationService.AddNotificationRestriction(notificationId, restriction)
	c.JSON(200, gin.H{"message": "restriction added"})
}
