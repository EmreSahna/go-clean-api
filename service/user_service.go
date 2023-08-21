package service

import (
	"go-clean-api/repository"
)

type UserServiceInterface interface {
}

type UserService struct {
	UserRepository *repository.UserRepository
}
