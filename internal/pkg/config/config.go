// Package config contains all the service configuration values.
// The configuration is from the config.yaml file.
package config

import (
	"github.com/billykore/go-service-tmpl/internal/pkg/config/internal"
	"github.com/billykore/go-service-tmpl/internal/pkg/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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
	logger := log.NewLogger()

	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("failed to read config file: %v", err)
	}

	cfg := new(Config)
	err := viper.Unmarshal(cfg)
	if err != nil {
		logger.ZapLogger().Info("failed to unmarshal config: %v", zap.Error(err))
	}

	return &cfg.Configs
}
