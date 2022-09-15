package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP     `json:"http"`
		LogLevel `json:"log"`
	}
	HTTP struct {
		HttpPort string `json:"http_port"`
	}
	LogLevel struct {
		Level string `json:"log_level"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config.json", cfg)
	if err != nil {
		return nil, err
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
