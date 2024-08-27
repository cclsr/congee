package config

type Config struct {
	Mysql  MysqlConfig  `yaml:"mysql"`
	Logger LogConfig    `yaml:"logger"`
	System SystemConfig `yaml:"system"`
}
