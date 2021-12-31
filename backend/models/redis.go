package models

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func OpenRedisConn() *redis.Client {
	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	pass := fmt.Sprintf("%s", os.Getenv("REDIS_PASS"))
	redis := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})
	return redis
}
