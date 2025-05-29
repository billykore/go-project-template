package postgres

import (
	"github.com/billykore/go-service-tmpl/internal/pkg/config"
	"github.com/billykore/go-service-tmpl/internal/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New returns new postgres db connection.
func New(cfg *config.Configs) *gorm.DB {
	dsn := cfg.Postgres.DSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.NewLogger().Fatalf("failed to connect database: %v", err)
	}
	return db
}
