package internal

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func newRedisClient(conf *Conf) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.RedisHost, conf.RedisPort),
		Password: conf.RedisPass,
		DB:       conf.RedisDB,
	})
	return rdb
}
