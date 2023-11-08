package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notifications/models"
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

func (h *UserHandlers) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, h.userService.GetUsers())
}

func (h *UserHandlers) GetUserById(c *gin.Context) {
	userId := c.Param("userId")
	user := h.userService.GetUserById(userId)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandlers) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.userService.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}

func (h *UserHandlers) DeleteUserById(c *gin.Context) {
	userId := c.Param("userId")
	h.userService.DeleteUserById(userId)
	c.JSON(200, gin.H{"message": "user deleted"})
}

func (h *UserHandlers) UserNotify(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user_id is empty"})
	}

	var userNotification models.UserNotifyRequest
	if err := c.ShouldBindJSON(&userNotification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.UserNotify(userId, userNotification); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user notified"})
}
