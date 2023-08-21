package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-clean-api/config"
	"go-clean-api/database"
	"go-clean-api/middleware"
	"go-clean-api/router"
	"go-clean-api/service"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	cfg := config.LoadConfig()

	r := gin.New()
	r.Use(middleware.AllowCors())

	db := database.ConnDB(cfg)

	services := service.NewServices(db, cfg)

	router.SetupRoutes(r, services)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Info("Starting server", zap.Int("port", cfg.Port), zap.String("srv", srv.Addr))
	err := srv.ListenAndServe()
	logger.Fatal("Failed to start server", zap.Error(err))
}
