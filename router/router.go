package router

import (
	"github.com/gin-gonic/gin"
	"go-clean-api/handler"
	"go-clean-api/service"
)

func SetupRoutes(engine *gin.Engine, services *service.Services) {
	// Healthcheck
	healthCheckHandler := handler.NewHealthCheckHandler(services.HealthCheckService)
	health := engine.Group("/healthcheck")
	{
		health.GET("", healthCheckHandler.Healthcheck)
	}

	// Users
	userHandler := handler.NewUserHandler(services.UserService)
	user := engine.Group("/user")
	{
		user.POST("", userHandler.CreateUser)
		user.GET("", userHandler.GetUsers)
		user.GET("/:id", userHandler.GetUser)
	}
}
