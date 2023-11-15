package config

type RedisBaseConfig struct {
	Timeout     int `yaml:"Timeout" toml:"Timeout"` // dial timeout `yaml:"name"`
	MinIdle     int `yaml:"MinIdle" toml:"MinIdle"`
	IdleTimeout int `yaml:"IdleTimeout" toml:"IdleTimeout"`
	PoolSize    int `yaml:"PoolSize" toml:"PoolSize"`
	MaxActive   int
}

// RedisConfig 单实例Redis配置
type RedisConfig struct {
	Addr     string `yaml:"Addr" toml:"Addr"` // host:port
	Password string `yaml:"Password" toml:"Password"`
	RedisBaseConfig
}

// RedisGroupConfig Redis组配置
type RedisGroupConfig struct {
	Addrs     []string `yaml:"Addrs" toml:"Addrs"`
	Passwords []string `yaml:"Passwords" toml:"Passwords"`
	RedisBaseConfig
}

// RedisClusterConfig Redis集群配置
type RedisClusterConfig struct {
	Addr     []string `yaml:"Addr" toml:"Addr"`
	Password string   `yaml:"Password" toml:"Password"`
	ReadOnly bool     `yaml:"ReadOnly" toml:"ReadOnly"`
	RedisBaseConfig
}
