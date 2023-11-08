package routes

import (
	"github.com/gin-gonic/gin"
	"notifications/handlers"
)

func SetupUserRoutes(r *gin.Engine, userHandler *handlers.UserHandlers) {
	//CRUD
	r.GET("/users", userHandler.GetUsers)
	r.GET("/user/:userId", userHandler.GetUserById)
	r.POST("/user", userHandler.CreateUser)
	r.DELETE("/user/:userId", userHandler.DeleteUserById)

	//Notity user
	r.POST("/user/:userId/notify", userHandler.UserNotify)
}
