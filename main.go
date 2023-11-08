package main

import (
	"github.com/gin-gonic/gin"
	"notifications/handlers"
	"notifications/repositories"
	"notifications/routes"
	"notifications/services"
)

func main() {
	r := gin.Default()

	notificationRepository := repositories.NewNotificationRepository()
	userRepository := repositories.NewUserRepository()

	userService := services.NewUserService(userRepository, notificationRepository)
	userHandler := handlers.NewUserHandlers(userService)
	routes.SetupUserRoutes(r, userHandler)

	notificationService := services.NewNotificationService(notificationRepository)
	notificationHandler := handlers.NewNotificationHandlers(notificationService)
	routes.SetupNotificationRoutes(r, notificationHandler)

	r.Run(":8080")
}
