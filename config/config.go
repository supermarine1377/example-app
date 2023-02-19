package config

import "github.com/caarlos0/env/v7"

type Config struct {
	Port string `env:"PORT" envDefault:"80"`
	Env  string `env:"ENV" envDefault:"dev"`
}

func New() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
