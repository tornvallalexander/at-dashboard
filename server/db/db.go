package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Database struct {
	Client *redis.Client
}

var ctx = context.Background()

func ConnectDB() (*Database, error) {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatalf("Failed to load env vars: %s\n", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CLIENT_ADDR"),
		Password: os.Getenv("REDIS_CLIENT_PASSWORD"),
		DB:       0,
	})

	if err = rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Database{
		Client: rdb,
	}, nil
}
