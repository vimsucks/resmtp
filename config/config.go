package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Mode string
	Listen string
	Port int
	Database struct {
		URL string
		Type string
	}
}

var conf Config

func InitConfig(cliConf CliConfig) Config {
	if _, err := toml.DecodeFile(cliConf.ConfigPath, &conf); err != nil {
		panic(err)
	}
	return conf
}

func GetGlobalConfig() Config {
	return conf
}
