package repositories

import (
	"os"

	redis "github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func createRedisConnection() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return redisClient

}
