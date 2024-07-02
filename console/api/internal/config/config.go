package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	AuthCfg struct {
		AccessSecret string
		AccessExpire int64
	}
	RedisConf redis.RedisConf
	Mysql     struct {
		DataSource string
		MaxIdle    int
		MaxOpen    int
		Level      string
		DbCache    cache.CacheConf
	}
}
