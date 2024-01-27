package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/v7ktory/htmx+go/internal/handler"
	"github.com/v7ktory/htmx+go/internal/service"
	"github.com/v7ktory/htmx+go/pkg/database/postgres"
	postgresdb "github.com/v7ktory/htmx+go/pkg/database/postgres/sqlc"
	"github.com/v7ktory/htmx+go/pkg/database/redis"
	"github.com/v7ktory/htmx+go/pkg/session"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redis, err := redis.NewRedisDB(redis.Config{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       os.Getenv("REDIS_DB"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	postgres, err := postgres.NewPostgresDB(postgres.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB_NAME"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	q := postgresdb.New(postgres)

	sessionManager := session.NewSessionManager(redis)

	service := service.NewService(redis, q)
	handler := handler.NewHandler(sessionManager, service)

	router := handler.InitRoute()

	log.Fatal(router.Run(":3000"))
}
