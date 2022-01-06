package common

import "github.com/gomodule/redigo/redis"

func InitRedis(addr string, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     2,
		IdleTimeout: 240,
		MaxActive:   3,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}

}
