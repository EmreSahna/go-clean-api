package handler

import (
	"gin-example/pkg"
	"gin-example/pkg/models"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	usersService pkg.Users
}

func NewUsersHandler(handler *gin.Engine, us pkg.Users) {
	h := &UsersHandler{us}
	handler.POST("/users", h.createUser)
}

func (h *UsersHandler) createUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := h.usersService.CreateUser(user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User created"})
}
