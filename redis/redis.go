package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Connect(uri string, user string, password string) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     uri,
		Username: user,
		Password: password,
		DB:       0, // use default DB
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	return nil
}
