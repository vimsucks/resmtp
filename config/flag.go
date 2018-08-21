package config

import "flag"

type CliConfig struct {
	ConfigPath string
}

func ParseCli() CliConfig {
	cliConf := CliConfig{}
	flag.StringVar(&cliConf.ConfigPath, "conf", "config.toml", "config file path")
	flag.Parse()
	return cliConf
}
