package redis

import (
	"context"

	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	redis "github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

func init() {
	Client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := Client.Ping(ctx).Result()
	if err != nil {
		logger.Error("cannot ping redis", err)
		panic(err)
	}

	logger.Info("redis successfuly connected")
}
