package routes

import (
	"github.com/gin-gonic/gin"
	"notifications/handlers"
)

func SetupNotificationRoutes(r *gin.Engine, notificationHandler *handlers.NotificationHandlers) {
	//CRUD
	r.GET("/notifications", notificationHandler.GetNotifications)
	r.GET("/notification/:notificationId", notificationHandler.GetNotificationById)
	r.POST("/notification", notificationHandler.CreateNotification)
	r.DELETE("/notification/:notificationId", notificationHandler.DeleteNotificationById)

	//Notity notification
	r.POST("/notification/:notificationId/restriction", notificationHandler.NotificationRestriction)
}
