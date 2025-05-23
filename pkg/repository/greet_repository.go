package repository

import (
	"context"

	"github.com/billykore/go-service-tmpl/pkg/entity"
)

// GreetRepository defines the methods to interacting with persistence storage used by service.
type GreetRepository interface {
	// GetAll retrieves all Greet entities from the repository. Returns an error if retrieval fails.
	GetAll(ctx context.Context) ([]entity.Greet, error)

	// Save stores a Greet in the persistence storage. Returns an error if saving fails.
	Save(ctx context.Context, greet entity.Greet) error
}
