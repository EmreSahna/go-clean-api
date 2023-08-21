package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"go-clean-api/config"
	"log"
	"time"
)

func ConnDB(cfg *config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.Db.Dsn)

	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Failed to ping db", err)
	}

	return db
}
