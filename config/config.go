package config

import (
	"github.com/BurntSushi/toml"
)

type ServConfig struct {
	BindAddr     string   `toml:"bindAddr"`
	ReadTimeout  int      `toml:"readTimeout"`
	WriteTimeout int      `toml:"writeTimeout"`
	AllowOrigins []string `toml:"allowOrigins"`
}

func NewConfig() *ServConfig {
	return &ServConfig{}
}

// TODOмб указатель использовать?
func ReadConfigFile(path string, dst interface{}) error {
	config := NewConfig()
	_, err := toml.DecodeFile(path, config)
	return err
}

// type DataBaseConfigT struct {
// 	User     string `toml:"user"`
// 	Password string `toml:"password"`
// 	Host     string `toml:"host"`
// 	Port     int    `toml:"port"`
// 	DBname   string `toml:"dbname"`
// }

// func (c *ServConfig) GetPostgresConfig() string {
// 	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
// 		c.DataBaseConfig.Host, c.DataBaseConfig.Port, c.DataBaseConfig.User, c.DataBaseConfig.Password, c.DataBaseConfig.DBname)
// }
