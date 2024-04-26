package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	LogFile *LogFileConfig
	JWT     *JWTConfig
}

type LogFileConfig struct {
	AccessLogPath string
	ErrorLogPath  string
	MaxSize       int
}

type JWTConfig struct {
	SecretKey string
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
