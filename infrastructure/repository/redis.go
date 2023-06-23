package repositories

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func createRedisConnection() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	})
	return redisClient

}
