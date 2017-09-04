package util

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// InitRedis ...
func InitRedis() (Pool *redis.Pool, err error) {
	Pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", Config.RedisConfig["core"])
			return c, err
		},
	}
	return Pool, err
}
