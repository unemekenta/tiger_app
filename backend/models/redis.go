package models

import (
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

func RedisNewPool() *redis.Pool {
	redisUrl := os.Getenv("REDIS_URL")
	tsl, _ := strconv.ParseBool(os.Getenv("REDIS_TSL"))

	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   0,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisUrl, redis.DialTLSSkipVerify(tsl))
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}
