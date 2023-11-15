package config

type IConfig interface {
	SchemeConfig() map[string]*SchemeConfig
	RedisConfig() map[string]*RedisConfig
	RedisGroupConfig() map[string]*RedisGroupConfig
	RedisClusterConfig() map[string]*RedisClusterConfig
	KafkaConfig() map[string]*KafkaConfig
	ClickHouseConfig() map[string]*ClickHouseConfig
}

type Config struct {
	Scheme       map[string]*SchemeConfig       `yaml:"Scheme",toml:"Scheme"`
	ClickHouse   map[string]*ClickHouseConfig   `yaml:"ClickHouse",toml:"ClickHouse"`
	RedisGroup   map[string]*RedisGroupConfig   `yaml:"RedisGroup",toml:"RedisGroup"`
	Redis        map[string]*RedisConfig        `yaml:"Redis",toml:"Redis"`
	RedisCluster map[string]*RedisClusterConfig `yaml:"RedisCluster",toml:"RedisCluster"`
	Kafka        map[string]*KafkaConfig        `yaml:"Kafka",toml:"Kafka"`
}

type CronAppConfig struct {
	Config
}

// SchemeConfig 获取关系型数据库配置，如Mysql
func (c *CronAppConfig) SchemeConfig() map[string]*SchemeConfig {
	return c.Scheme
}

// ClickHouseConfig 获取Redis组配置
func (c *CronAppConfig) ClickHouseConfig() map[string]*ClickHouseConfig {
	return c.ClickHouse
}

// RedisConfig 获取Redis配置
func (c *CronAppConfig) RedisConfig() map[string]*RedisConfig {
	return c.Redis
}

// RedisGroupConfig 获取Redis组配置
func (c *CronAppConfig) RedisGroupConfig() map[string]*RedisGroupConfig {
	return c.RedisGroup
}

// RedisClusterConfig 获取集群Redis配置
func (c *CronAppConfig) RedisClusterConfig() map[string]*RedisClusterConfig {
	return c.RedisCluster
}

// KafkaConfig 获取Kafka配置
func (c *CronAppConfig) KafkaConfig() map[string]*KafkaConfig {
	return c.Kafka
}

var _ IConfig = &CronAppConfig{}
