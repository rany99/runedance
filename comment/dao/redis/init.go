package redis

import (
	"github.com/gomodule/redigo/redis"
	"runedance/common/config"
	"runedance/common/util"
)

var redisPool *redis.Pool
var expireTimeUtil util.ExpireTimeUtil

func Init() {
	redisPool = &redis.Pool{
		MaxIdle:   config.Redis.MaxIdle,
		MaxActive: config.Redis.MaxActive,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.Redis.Address)
		},
	}
	expireTimeUtil = util.ExpireTimeUtil{
		ExpireTime:     config.Redis.ExpireTime,
		MaxRandAddTime: config.Redis.MaxRandAddTime,
	}

}
