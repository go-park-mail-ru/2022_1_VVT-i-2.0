package serv

import (
	"go.uber.org/zap"
)

type Config struct {
	BindAddr     string     `toml:"bindAddr"`
	ReadTimeout  int        `toml:"readTimeout"`
	WriteTimeout int        `toml:"writeTimeout"`
	LogConfig    zap.Config `toml:"logger"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
