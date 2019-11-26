package helper

import (
	"github.com/gomodule/redigo/redis"
	"sync"
)

var (
	//ctx context.Context
	err       error
	once      sync.Once
	RedisConn *redis.Pool
)

func newRedis() {
	RedisConn = &redis.Pool{
		Dial:            nil,
		TestOnBorrow:    nil,
		MaxIdle:         0,
		MaxActive:       0,
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}

}
