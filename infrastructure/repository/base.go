package repositories

import (
	"sync"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type RedisDB struct {
	Client *redis.Client
}

type PGRepository struct {
	DB *gorm.DB
}

func RedisInit() *RedisDB {
	o := sync.Once{}
	var db *RedisDB
	o.Do(func() {
		db = &RedisDB{
			Client: createRedisConnection(),
		}
	})
	return db
}

func NewGormDatabase() *PGRepository {
	db, _ := GormInit()
	return &PGRepository{DB: db}
}
