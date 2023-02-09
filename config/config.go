package config

import (
	"github.com/spf13/viper"
	"go-transactions-gateway/pkg/postgres"
	"log"
)

type (
	Config struct {
		Logger   Logger          `mapstructure:"logger" validate:"required"`
		Server   Server          `mapstructure:"server" validate:"required"`
		Postgres postgres.Config `mapstructure:"postgres" validate:"required"`
	}

	Logger struct {
		Level *int8 `mapstructure:"level" validate:"required"`
	}

	Server struct {
		Host string `mapstructure:"host" validate:"required"`
		Port int    `mapstructure:"port" validate:"required"`
	}
)

func Load() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("config")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("config reading error: %v", err)
	}

	cfg := &Config{}

	err = v.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
