package postgres

import (
	"github.com/billykore/go-service-tmpl/internal/pkg/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New returns new postgres db connection.
func New(cfg *config.Configs) *gorm.DB {
	dsn := cfg.Postgres.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect database")
	}
	return db
}
