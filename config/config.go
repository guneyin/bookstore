package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	defaultPort = 8080
)

type Config struct {
	Port int `env:"PORT"`
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
