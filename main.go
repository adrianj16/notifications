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

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandlers(userService)
	routes.SetupUserRoutes(r, userHandler)




	r.Run(":8080")
}
