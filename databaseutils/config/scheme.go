package config

// SchemeConfig 关系型数据库的配置 mysql pgsql mssql
type SchemeConfig struct {
	// mysql pgsql
	Driver string `yaml:"Driver" toml:"driver"`
	// dsn example
	// param参数设置参考 https://github.com/go-sql-driver/mysql#parameters
	// username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&readTimeout=3s&timeout=3s&writeTimeout=3s&rejectReadOnly=false&parseTime=true&loc=Local
	// param 参数：
	//
	Dsn string `yaml:"Dsn" toml:"dsn"`
	// 建立最大的连接数
	MaxOpenConns int `yaml:"MaxOpenConns" toml:"maxOpenConns"`
	// 最大的空闲连接数
	MaxIdleConns int `yaml:"MaxIdleConns" toml:"maxIdleConns"`
	// 一条连接存活的最长时间
	ConnMaxLifeTime int `yaml:"ConnMaxLifeTime" toml:"connMaxLifeTime"`
}
