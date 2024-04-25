package config

type Config struct {
	LogFile *LogFileConfig
}

type LogFileConfig struct {
	Path    string
	MaxSize int
}

var Cfg *Config

func init() {

}
