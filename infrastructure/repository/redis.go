package repositories

import (
	

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func createRedisConnection() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return redisClient

}
