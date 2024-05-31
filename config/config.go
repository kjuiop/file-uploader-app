package config

import "github.com/kelseyhightower/envconfig"

type EnvConfig struct {
	Server Server
}

type Server struct {
	Mode string `envconfig:"FUA_ENV" default:"dev"`
	Port string `envconfig:"FUA_SERVER_PORT" default:"8090"`
}

func LoadEnvConfig() (*EnvConfig, error) {
	var config EnvConfig
	err := envconfig.Process("fua", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
