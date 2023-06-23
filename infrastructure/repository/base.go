package repositories



import (
	"sync"

	"github.com/go-redis/redis/v8"
	
)

type RedisDB struct {
	client *redis.Client
}

func RedisInit() *RedisDB {
	o := sync.Once{}
	var db *RedisDB
	o.Do(func() {
		db = &RedisDB{
			client: createRedisConnection(),
		}
	})
	return db
}

