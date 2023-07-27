package pkg

import "gin-example/pkg/models"

type GreetingsRepo interface {
	GetGreeting() string
}

type Greetings interface {
	GetGreeting(name string) string
	GetGreetingFull(greetings models.Greetings) string
}

type UsersRepo interface {
	CreateUser(user models.User) error
}

type Users interface {
	CreateUser(user models.User) error
}
