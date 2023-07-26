package services

import (
	"gin-example/pkg"
	"gin-example/pkg/models"
)

type GreetingsService struct {
	repo pkg.GreetingsRepo
}

func NewGreetingsService(repo pkg.GreetingsRepo) *GreetingsService {
	return &GreetingsService{repo: repo}
}

func (s *GreetingsService) GetGreeting(name string) string {
	if name == "" {
		name = "stranger"
	}
	return s.repo.GetGreeting() + " " + name
}

func (s *GreetingsService) GetGreetingFull(greetings models.Greetings) string {
	if greetings.Name == "" {
		greetings.Name = "stranger"
	}
	if greetings.Surname == "" {
		greetings.Surname = "stranger"
	}
	return s.repo.GetGreeting() + " " + greetings.Name + " " + greetings.Surname
}
