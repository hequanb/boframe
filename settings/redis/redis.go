package redis

import (
	"context"
	"fmt"
	
	"boframe/settings"
	"github.com/go-redis/redis/v8"
)

var redisCli *redis.Client

func Init(config *settings.RedisConfig) (err error) {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})
	
	_, err = redisCli.Ping(context.Background()).Result()
	return err
}

func Close() {
	redisCli.Close()
}
