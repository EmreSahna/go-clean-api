package main

import (
	"fmt"
	"gin-example/api/handler"
	"gin-example/pkg/repository"
	"gin-example/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	greetingsService := services.NewGreetingsService(
		repository.NewGreetingsRepository(),
	)

	r := gin.New()
	handler.NewGreetingsHandler(r, greetingsService)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
}
