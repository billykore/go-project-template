// Package config contains all the service configuration values.
// The configuration is from the config.yaml file.
package config

import (
	"github.com/billykore/go-service-tmpl/internal/pkg/config/internal"
	"github.com/billykore/go-service-tmpl/internal/pkg/log"
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
	logger := log.New()

	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("failed to read config file: %v", err)
	}

	cfg := new(Config)
	err := viper.Unmarshal(cfg)
	if err != nil {
		logger.Infof("failed to unmarshal config: %v", err)
	}

	return &cfg.Configs
}
