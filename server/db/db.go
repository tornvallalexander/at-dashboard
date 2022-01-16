package db

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

var ctx = context.Background()

func ConnectDB() (*Database, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-13927.c250.eu-central-1-1.ec2.cloud.redislabs.com:13927",
		Password: "gbOj6hYu83JaKFHxB8e5QUpB40gTgJXR",
		DB:       0,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Database{
		Client: rdb,
	}, nil
}
