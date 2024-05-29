package config

import "github.com/kelseyhightower/envconfig"

type EnvConfig struct {
	Server Server
}

type Server struct {
	IsDryRun bool   `envconfig:"FUA_DRY_RUN" default:"true"`
	Port     string `envconfig:"FUA_SERVER_PORT" default:"8090"`
}

func LoadEnvConfig() (*EnvConfig, error) {
	var config EnvConfig
	err := envconfig.Process("fua", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
