package handler

import (
	"github.com/gin-gonic/gin"
	"go-clean-api/responses"
	"go-clean-api/service"
)

type UserHandler struct {
	UserService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var response responses.Response
	c.JSON(response.Status, response.Body)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	var response responses.Response
	c.JSON(response.Status, response.Body)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	var response responses.Response
	c.JSON(response.Status, response.Body)
}
