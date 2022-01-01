package models

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

func RedisNewPool() *redis.Pool {
	redisUrl := fmt.Sprintf("redis://%s@%s:%s", os.Getenv("REDIS_PASS"), os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	tsl, _ := strconv.ParseBool(os.Getenv("REDIS_PASS"))

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
