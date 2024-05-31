package config

import "github.com/kelseyhightower/envconfig"

type EnvConfig struct {
	Server Server
	Logger Logger
}

type Server struct {
	Mode string `envconfig:"FUA_ENV" default:"dev"`
	Port string `envconfig:"FUA_SERVER_PORT" default:"8090"`
}

type Logger struct {
	Level       string `envconfig:"AWC_LOG_LEVEL" default:"debug"`
	Path        string `envconfig:"AWC_LOG_PATH" default:"./logs/access.log"`
	PrintStdOut bool   `envconfig:"LOG_STDOUT" default:"false"`
}

func LoadEnvConfig() (*EnvConfig, error) {
	var config EnvConfig
	err := envconfig.Process("fua", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
