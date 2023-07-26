package handler

import (
	"gin-example/api/responses"
	"gin-example/pkg"
	"gin-example/pkg/models"
	"github.com/gin-gonic/gin"
)

type GreetingsHandler struct {
	greetingsService pkg.Greetings
}

func NewGreetingsHandler(handler *gin.Engine, gs pkg.Greetings) {
	h := &GreetingsHandler{gs}
	handler.GET("/greetings", h.getGreeting)
	handler.POST("/greetings/full", h.getGreetingFull)
}

func (h *GreetingsHandler) getGreeting(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, gin.H{
		"greeting": h.greetingsService.GetGreeting(name),
	})
}

func (h *GreetingsHandler) getGreetingFull(c *gin.Context) {
	var greetings models.Greetings
	if err := c.ShouldBindJSON(&greetings); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}

	res := responses.GreetingsResponse{
		Greeting: h.greetingsService.GetGreetingFull(greetings),
	}

	c.JSON(200, res)
}
