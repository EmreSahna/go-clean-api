package main

import (
	"database/sql"
	"fmt"
	"gin-example/api/handler"
	"gin-example/pkg/repository"
	"gin-example/pkg/services"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		panic(err)
	}

	greetingsService := services.NewGreetingsService(
		repository.NewGreetingsRepository(),
	)

	r := gin.New()
	handler.NewGreetingsHandler(r, greetingsService)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
}
