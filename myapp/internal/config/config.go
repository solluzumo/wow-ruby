package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port        string `env:"PORT" envDefault:"8080"`
	JWTSecret   string `env:"JWT_SECRET" envDefault:"secret"`
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
