package config

type LogConfig struct {
	Level        string `yaml:"evel"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"in_console"`
}
