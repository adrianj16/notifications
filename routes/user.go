package routes

import (
	"github.com/gin-gonic/gin"
	"notifications/handlers"
)

func SetupUserRoutes(r *gin.Engine, userHandler *handlers.UserHandlers) {
	r.GET("/hello", userHandler.Register)
	r.GET("/login", userHandler.Login)
}
