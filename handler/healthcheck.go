package handler

import (
	"github.com/gin-gonic/gin"
	"go-clean-api/responses"
	"go-clean-api/service"
)

type HealthCheckHandler struct {
	HealthService service.HealthCheckServiceInterface
}

func NewHealthCheckHandler(healthService service.HealthCheckServiceInterface) *HealthCheckHandler {
	return &HealthCheckHandler{HealthService: healthService}
}

func (h *HealthCheckHandler) Healthcheck(c *gin.Context) {
	var response responses.Response
	h.HealthService.Healthcheck(&response)
	c.JSON(response.Status, response.Body)
}
