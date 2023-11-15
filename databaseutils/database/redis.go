package database

import (
	"github.com/HadesTso/databaseConnect/databaseutils/config"
	"github.com/go-redis/redis"
	"time"
)

func RedisConn(redisName string) *redis.Client {
	if g, ok := Container.Redis[redisName]; ok {
		return g
	} else {
		if l, ok := sourceLock.Redis[redisName]; !ok {
			return nil
		} else {
			l.Lock()
			defer l.Unlock()
			if c, ok := Container.Redis[redisName]; ok {
				return c
			}

			RedisConf := connConfig.RedisConfig()
			for rid, conf := range RedisConf {
				if rid == redisName {
					Container.Redis[rid] = newRedis(conf)
					return Container.Redis[rid]
				}
			}
		}
	}

	return nil
}

func newRedis(conf *config.RedisConfig) *redis.Client {
	opt := &redis.Options{}
	opt.Addr = conf.Addr
	opt.DB = 0
	opt.DialTimeout = time.Duration(conf.Timeout) * time.Second
	opt.Password = conf.Password
	opt.PoolSize = conf.PoolSize
	opt.IdleTimeout = time.Duration(conf.IdleTimeout) * time.Second
	opt.MinIdleConns = conf.MinIdle
	return redis.NewClient(opt)
}
