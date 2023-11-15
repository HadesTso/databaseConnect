package config

type KafKaBaseConfig struct {
}

type Sasl struct {
	User      string `yaml:"User" toml:"User"`
	Password  string `yaml:"Password" toml:"Password"`
	Mechanism string `yaml:"Mechanism" toml:"Mechanism"`
}

type KafkaConfig struct {
	Addr     string `yaml:"Addr" toml:"Addr"` // host:port
	Version  string `yaml:"Version" toml:"Version"`
	Sasl     Sasl   `yaml:"Sasl" toml:"Sasl"`
	Topic    string `yaml:"Topic" toml:"Topic"`
	Earliest bool   `yaml:"Earliest" toml:"Earliest"`
}
