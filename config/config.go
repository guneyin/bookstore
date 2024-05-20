package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	defaultPort = 8080
)

type Config struct {
	Port           int    `env:"PORT"`
	SmtpPort       int    `env:"SMTP_PORT"`
	SmtpPassword   string `env:"SMTP_PWD"`
	SmtpUserName   string `env:"SMTP_USER"`
	SmtpServer     string `env:"SMTP_SERVER"`
	SenderEmail    string `env:"SMTP_SENDER"`
	SenderIdentity string `env:"SMTP_IDENTITY"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		return nil, err
	}

	err = cfg.validate()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) validate() error {
	switch {
	case c.Port == 0:
		c.Port = defaultPort
	}

	return nil
}
