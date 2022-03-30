package config

import (
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ServConfig   ServerConfig      `toml:"server"`
	LoggerConfig LogConfig         `toml:"logger"`
	AuthConfig   AuthManagerConfig `toml:"authManager"`
}

type ServerConfig struct {
	BindAddr     string   `toml:"bindAddr"`
	ReadTimeout  int      `toml:"readTimeout"`
	WriteTimeout int      `toml:"writeTimeout"`
	AllowOrigins []string `toml:"allowOrigins"`
}

type LogConfig struct {
	Level            string
	Encoding         string
	OutputPaths      []string
	ErrorOutputPaths []string

	MessageKey    string
	TimeKey       string
	LevelKey      string
	NameKey       string
	FunctionKey   string
	StacktraceKey string
}

type AuthManagerConfig struct {
	Key    string
	Method string
	// ExpiryTime time.Duration `toml:"expiryTime"`
	ExpiryTime duration
}
type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func NewConfig() *Config {
	return &Config{}
}

func ReadConfigFile(configPath string, dst interface{}) error {
	_, err := toml.DecodeFile(configPath, dst)
	return err
}
