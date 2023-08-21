package service

import (
	"database/sql"
	"go-clean-api/config"
	"go-clean-api/repository"
)

type Services struct {
	UserService        *UserService
	HealthCheckService *HealthCheckService
}

func NewServices(db *sql.DB, cfg *config.Config) *Services {
	healthCheckService := &HealthCheckService{cfg}

	repositories := repository.NewRepositories(db)
	userService := &UserService{UserRepository: repositories.UserRepository}

	return &Services{
		UserService:        userService,
		HealthCheckService: healthCheckService,
	}
}
