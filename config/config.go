package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	LogFile *LogFileConfig
	JWT     *JWTConfig
	MySQL   *MySQLConfig
}

type LogFileConfig struct {
	AccessLogPath string
	ErrorLogPath  string
	MaxSize       int
}

type JWTConfig struct {
	SecretKey string
}

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     int64
	DB       string
}

var Cfg *Config

func init() {
	var err error
	viper.SetConfigFile("./config.yaml")
	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	}

	if err = viper.Unmarshal(&Cfg); err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	}
}
