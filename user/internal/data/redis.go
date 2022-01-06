package data

import (
	"github.com/gomodule/redigo/redis"
	"github.com/tjl-cmd/travel/common"
	"github.com/tjl-cmd/travel/user/internal/conf"
)

var (
	Redispool *redis.Pool
)

func InitRedisPool(bc *conf.Bootstrap) {
	Redispool = common.InitRedis(bc.Data.Redis.Addr, bc.Data.Redis.Password)
}
