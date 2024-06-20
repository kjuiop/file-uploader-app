package config

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	Server Server
	Logger Logger
	Slack  Slack
	Upload Upload
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

type Slack struct {
	WebhookReportUrl string `envconfig:"FUA_SLACK_WEBHOOK_REPORT_URL" `
}

type Upload struct {
	FileSourceDir string `envconfig:"FUA_FILE_SOURCE_DIR" default:"/home/jake/file-upload"`
}

func LoadEnvConfig() (*EnvConfig, error) {

	var config EnvConfig
	if err := envconfig.Process("fua", &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *EnvConfig) CheckValid() error {
	if c.Slack.WebhookReportUrl == "" {
		return errors.New("webhook report url required")
	}

	return nil
}
