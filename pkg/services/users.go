package services

import (
	"gin-example/pkg"
	"gin-example/pkg/models"
)

type UsersService struct {
	repo pkg.UsersRepo
}

func NewUsersService(repo pkg.UsersRepo) *UsersService {
	return &UsersService{repo}
}

func (s *UsersService) CreateUser(user models.User) error {
	return s.repo.CreateUser(user)
}
