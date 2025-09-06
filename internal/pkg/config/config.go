// Package config contains all the service configuration values.
// The configuration is from the config.yaml file.
package config

import (
	"github.com/billykore/go-service-tmpl/internal/pkg/config/internal"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Configs hold the application configurations.
type Configs struct {
	App      internal.App
	Postgres internal.Postgres
	Token    internal.Token
}

type Config struct {
	Name    string
	Version string
	Configs Configs
}

// Load loads application configuration from a YAML file using Viper.
func Load() *Configs {
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Panic().Err(err).Msg("failed to read config file")
	}

	cfg := new(Config)
	err := viper.Unmarshal(cfg)
	if err != nil {
		log.Panic().Err(err).Msg("failed to unmarshal config")
	}

	return &cfg.Configs
}
