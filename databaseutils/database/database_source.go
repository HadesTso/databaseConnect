package database

import (
	"errors"
	"github.com/HadesTso/databaseConnect/databaseutils/config"
	"github.com/IBM/sarama"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	gormIo "gorm.io/gorm"
	"log"
	"sync"
)

type SourceContainer struct {
	DB                 map[string]*gorm.DB
	DBIo               map[string]*gormIo.DB
	Clickhouse         map[string]*gorm.DB
	RedisGroup         map[string]*RedisGroupContainer
	Redis              map[string]*redis.Client
	RedisCluster       map[string]*redis.ClusterClient
	KafkaAsyncProducer map[string]sarama.AsyncProducer
	KafkaConsumer      map[string]*KafkaConsumer
}

type sourceContainerLock struct {
	DB                 map[string]*sync.Mutex
	DBIo               map[string]*sync.Mutex
	Clickhouse         map[string]*sync.Mutex
	RedisGroup         map[string]*sync.Mutex
	Redis              map[string]*sync.Mutex
	RedisCluster       map[string]*sync.Mutex
	KafkaAsyncProducer map[string]*sync.Mutex
	KafkaConsumer      map[string]*sync.Mutex
}

var (
	initSuccess bool // 初始化状态的值
	Container   *SourceContainer
	sourceLock  *sourceContainerLock
	connConfig  config.IConfig
)

func Init(config config.IConfig) *SourceContainer {
	log.Println("Db init start...")
	if initSuccess {
		panic(errors.New("denny again init"))
	}
	Container = &SourceContainer{}
	sourceLock = &sourceContainerLock{}

	redisConf := config.RedisConfig()
	if len(redisConf) > 0 {
		log.Println("redis init start...")
		Container.Redis = make(map[string]*redis.Client, len(redisConf))
		sourceLock.Redis = make(map[string]*sync.Mutex, len(redisConf))
		for rid, _ := range redisConf {
			sourceLock.Redis[rid] = &sync.Mutex{}
		}
	}

	redisGroupConf := config.RedisGroupConfig()
	if len(redisGroupConf) > 0 {
		log.Println("redisGroup init start...")
		Container.RedisGroup = make(map[string]*RedisGroupContainer, len(redisGroupConf))
		sourceLock.RedisGroup = make(map[string]*sync.Mutex, len(redisGroupConf))
		for rid, confs := range redisGroupConf {
			sourceLock.RedisGroup[rid] = &sync.Mutex{}
			Container.RedisGroup[rid] = &RedisGroupContainer{}
			Container.RedisGroup[rid].C = make([]*redis.Client, len(confs.Addrs))
			Container.RedisGroup[rid].Cnt = len(confs.Addrs)
			Container.RedisGroup[rid].Id = rid
		}
	}

	return Container
}
