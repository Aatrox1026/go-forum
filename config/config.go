package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	LogFile   *LogFileConfig
	JWT       *JWTConfig
	Snowflake *SnowflakeConfig
	MySQL     *MySQLConfig
	Redis     *RedisConfig
}

type LogFileConfig struct {
	AccessLogPath string
	ErrorLogPath  string
	MaxSize       int
}

type JWTConfig struct {
	SecretKey string
}

type SnowflakeConfig struct {
	MachineID int64
}

type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

var Cfg *Config

func Init() {
	var err error
	viper.SetConfigFile("./config.yaml")
	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	}

	if err = viper.Unmarshal(&Cfg); err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	}
}
