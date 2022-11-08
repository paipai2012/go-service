package engine

import (
	"sale-service/util"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisEngine struct {
	*redis.Client
}

var redisEngine *RedisEngine

var redisOnce sync.Once

func GetRedisEngine() *RedisEngine {
	return redisEngine
}

func NewRedisEngine() {
	appInfo := util.GetYamlConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr:         appInfo.Redis.Host,
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	redisOnce.Do(func() {
		rdh := new(RedisEngine)
		rdh.Client = rdb
		redisEngine = rdh
	})
}
