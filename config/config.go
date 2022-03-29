package config

// import (
// 	"fmt"

// 	"github.com/BurntSushi/toml"
// )

// type ServConfig struct {
// 	BindAddr     string   `toml:"server.bindAddr"`
// 	ReadTimeout  int      `toml:"readTimeout"`
// 	WriteTimeout int      `toml:"writeTimeout"`
// 	AllowOrigins []string `toml:"allowOrigins"`
// }

// // 1. можно раскидать все конфиги по разным файлам
// // 2. можно создать все структуры конфигов здесь и передавать их в создаватели объектов, которые должны будут парсить
// // 3. можно парсить прям здесь делать из map-ы структуру типа dst

// // func NewConfig() *ServConfig {
// // 	return &ServConfig{}
// // }

// // TODOмб указатель использовать?

// const (
// 	SERVER_CONFIG = iota
// 	LOGGER_CONFIG
// 	DATABASE_CONFIG
// 	AUTH_MANAGER_CONFIG
// )

// // servConfig := serv.NewConfig()
// // _, err := toml.DecodeFile(*configPath, servConfig)
// // if err != nil {
// // 	log.Fatal(err)
// // }

// func NewServConfig() *ServConfig {
// 	return &ServConfig{
// 		BindAddr: ":8080",
// 	}
// }

// func ReadConfigFile(configPath string, dst interface{}, configType int) error {
// 	// var dstTmp interface{}
// 	switch configType {
// 	case SERVER_CONFIG:
// 		{
// 			// servConfig := NewConfig()
// 			// _, err := toml.DecodeFile(configPath, servConfig)
// 			// dst = servConfig
// 			// return err

// 			// type tmpServConfig struct {
// 			// Server interface{} `toml:"server"`
// 			// }
// 			//--------
// 			dstTmp := &struct {
// 				Server interface{} `toml:"server"`
// 			}{Server: dst}

// 			_, err := toml.DecodeFile(configPath, dstTmp)
// 			dst = dstTmp.Server
// 			return err
// 		}
// 	// case LOGGER_CONFIG:
// 	// 	{
// 	// 		dstTmp = struct {
// 	// 			LoggerCongig interface{} `toml:"logger"`
// 	// 		}{
// 	// 			LoggerCongig: dst,
// 	// 		}
// 	// 		_, err := toml.DecodeFile(configPath, dstTmp)
// 	// 		return err
// 	// 	}
// 	// case DATABASE_CONFIG:
// 	// 	{
// 	// 		dstTmp = struct {
// 	// 			DbConfig interface{} `toml:"database"`
// 	// 		}{
// 	// 			DbConfig: dst,
// 	// 		}
// 	// 		_, err := toml.DecodeFile(configPath, dstTmp)
// 	// 		return err
// 	// 	}
// 	// case AUTH_MANAGER_CONFIG:
// 	// 	{
// 	// 		dstTmp = struct {
// 	// 			AuthManagerConfig interface{} `toml:"authManager"`
// 	// 		}{
// 	// 			AuthManagerConfig: dst,
// 	// 		}
// 	// 		_, err := toml.DecodeFile(configPath, dstTmp)
// 	// 		return err
// 	// 	}
// 	default:
// 		return fmt.Errorf("unknown config-type argument: %v", configType)
// 	}
// 	// _, err := toml.DecodeFile(configPath, dstTmp)
// 	// return err
// }

// // func ReadConfigFile(path string, dst interface{}, tag string) error {
// // 	// config := NewConfig()
// // 	// str :=
// // 	var i struct {
// // 		field interface{} `toml:"`
// // 	}
// // 	configFile, err := os.Open(path)
// // 	defer configFile.Close()
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	decoder := toml.NewDecoder(configFile)
// // 	decoder.SetTagName(path)
// // 	_, err := toml.DecodeFile(path, dst)
// // 	return err
// // }

// // type DataBaseConfigT struct {
// // 	User     string `toml:"user"`
// // 	Password string `toml:"password"`
// // 	Host     string `toml:"host"`
// // 	Port     int    `toml:"port"`
// // 	DBname   string `toml:"dbname"`
// // }

// // func (c *ServConfig) GetPostgresConfig() string {
// // 	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
// // 		c.DataBaseConfig.Host, c.DataBaseConfig.Port, c.DataBaseConfig.User, c.DataBaseConfig.Password, c.DataBaseConfig.DBname)
// // }
