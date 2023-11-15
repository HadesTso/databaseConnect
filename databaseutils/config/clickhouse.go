package config

type ClickHouseConfig struct {
	Driver          string `yaml:"Driver" toml:"driver"`                   // 驱动: mysql/sqlite3/oracle/mssql/postgres/clickhouse, 如果集群配置了驱动, 这里可以省略//tcp://host1:9000?username=user&password=qwerty&database=clicks&read_timeout=10&write_timeout=20&alt_hosts=host2:9000,host3:9000
	Dsn             string `yaml:"Dsn" toml:"dsn"`                         // 数据库链接
	SetMaxOpenConns int    `yaml:"SetMaxOpenConns" toml:"setMaxOpenConns"` // (连接池)最大打开的连接数，默认值为0表示不限制
	SetMaxIdleConns int    `yaml:"SetMaxIdleConns" toml:"setMaxIdleConns"` // (连接池)闲置的连接数, 默认0
	ConnMaxLifeTime int    `yaml:"ConnMaxLifeTime" toml:"connMaxLifeTime"` // (连接池)最长存活时间
	//Prefix          string `yaml:"Prefix",toml:"srefix"`                   //// 表前缀, 如果集群配置了前缀, 这里可以省略
}
