package handlers

import (
	"github.com/gin-gonic/gin"
	"notifications/services"
)

type UserHandlers struct {
	userService *services.UserService
}

func NewUserHandlers(userService *services.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) GetUserById(c *gin.Context) {
	userId := c.Param("userId")
	err := h.userService.GetUserById(userId)
	if err != nil {
		c.JSON(200, gin.H{"message": "message"})
	}
	c.JSON(200, gin.H{"message": "message"})
}

func (h *UserHandlers) Register(c *gin.Context) {
	err := h.userService.Register("")
	if err != nil {
		c.JSON(200, gin.H{"message": "message"})
	}
	c.JSON(200, gin.H{"message": "message"})
}
