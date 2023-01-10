package client

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v9"
)

var redisclient *redis.Client

func init() {
	redisclient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})

	if _, err := redisclient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	err := redisclient.Set(ctx, "test", "Welcome to Golang with Redis and MongoDB", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client connected successfully...")
}
func Redis() *redis.Client {
	return redisclient
}
