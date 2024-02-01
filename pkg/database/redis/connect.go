package redis

import (
	"context"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Addr     string
	Password string
	DB       string
}

func NewRedisDB(cfg Config) (*redis.Client, error) {

	DB, err := strconv.Atoi(cfg.DB)
	if err != nil {
		return nil, err
	}
	// Create a new Redis client with the specified options
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Println("Error connecting to redis:", err)
		return nil, err
	}
	return client, nil
}
