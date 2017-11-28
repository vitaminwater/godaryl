package kv

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/vitaminwater/daryl/config"
)

var Pool *redis.Pool

func Init() {
	url := config.AppContext.String("redis-url")
	Pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", url) },
	}
}
