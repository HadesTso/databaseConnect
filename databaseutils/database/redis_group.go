package database

import "github.com/go-redis/redis"

type RedisGroupContainer struct {
	C   []*redis.Client
	Cnt int
	Id  string
}
