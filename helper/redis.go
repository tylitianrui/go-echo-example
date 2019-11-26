package helper

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go-echo-example/pkg/setting"
)

var (
	redisConn *redis.Pool
)

func newRedis() {
	var (
		c   redis.Conn
		err error
	)
	redisConn = &redis.Pool{

		Dial: func() (conn redis.Conn, e error) {

			c, err = redis.Dial("tcp", fmt.Sprintf("%s:%d", setting.G_RedisConf.Host, setting.G_RedisConf.Port))
			//c, err = redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			if setting.G_RedisConf.Password != "" {
				if _, err = c.Do("AUTH", setting.G_RedisConf.Password); err != nil {
					c.Close()
					return nil, err
				}

			}
			return c, nil
		},

		MaxIdle:     setting.G_RedisConf.MaxIdle,
		MaxActive:   setting.G_RedisConf.MaxActive,
		IdleTimeout: setting.G_RedisConf.IdleTimeout,
	}

}

func SetupRedis() {
	newRedis()
}

// 获取redis连接
func GetRedisConn() redis.Conn {
	return redisConn.Get()
}

// 获取
func RedisGet(key string) (reply []byte, err error) {
	var (
		c redis.Conn
	)
	c = GetRedisConn()
	defer c.Close()

	if reply, err = redis.Bytes(c.Do("GET", key)); err != nil {
		return nil, err
	}
	return reply, nil

}

func RedisSet(key string, data interface{}, time int) (err error) {
	var (
		c redis.Conn
		d []byte
	)
	c = GetRedisConn()
	defer c.Close()

	if d, err = json.Marshal(data); err != nil {

		return err
	}

	if time <= 0 {
		_, err = c.Do("SET", key, string(d))
		return err
	}
	_, err = c.Do("SET", key, string(d), "EX", time)
	return err

}
