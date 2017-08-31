package util

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// InitRedis ...
func InitRedis() (Pool *redis.Pool, err error) {
	redisConnection := "localhost:6379"
	Pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConnection)
			return c, err
		},
	}
	return Pool, err
}
