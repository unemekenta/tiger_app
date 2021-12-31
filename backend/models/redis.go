package models

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func OpenRedisConn() *redis.Client {
	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	pass := fmt.Sprintf("%s", os.Getenv("REDIS_PASS"))
	insecureSkipVerify := true
	if os.Getenv("MODE") == "local" {
		insecureSkipVerify = false
	}
	redis := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: insecureSkipVerify,
		},
	})
	return redis
}
