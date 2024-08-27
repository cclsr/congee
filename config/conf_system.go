package config

type SystemConfig struct {
	host string `yaml:"host"`
	prot int    `yaml:"prot"`
	env  string `yaml:"env"`
}
