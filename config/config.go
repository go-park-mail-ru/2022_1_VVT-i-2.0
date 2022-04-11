package config

import (
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ServConfig            ServerConfig      `toml:"server"`
	LoggerConfig          LogConfig         `toml:"logger"`
	AuthentificatorConfig AuthManagerConfig `toml:"authManager"`
	NotificatorConfig     NotificatorConfig `toml:"notificator"`
	CacherConfig          CachConfig        `toml:"cacher"`
	DatabaseCongig        DatabaseConfig    `toml:"database"`
	CorsConfig            CorsConfig        `toml:"cors"`
	CsrfConfig            CsrfConfig        `toml:"csrf"`
}

type ServerConfig struct {
	BindAddr     string `toml:"bindAddr"`
	ReadTimeout  int    `toml:"readTimeout"`
	WriteTimeout int    `toml:"writeTimeout"`
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

type CachConfig struct {
	Host string
	Port int
}

type AuthManagerConfig struct {
	Key        string
	Method     string
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

type NotificatorConfig struct {
	ApiKey string
	Email  string
}

type DatabaseConfig struct {
	DbName        string
	User          string
	Password      string
	Port          int
	Host          string
	ConnectionMax int
}

type CorsConfig struct {
	AllowOrigins []string `toml:"allowOrigins"`
	MaxAge       int      `toml:"maxAge"`
}

type CsrfConfig struct {
	MaxAge int `toml:"maxAge"`
}

func NewConfig() *Config {
	return &Config{}
}

func ReadConfigFile(configPath string, dst interface{}) error {
	_, err := toml.DecodeFile(configPath, dst)
	return err
}
