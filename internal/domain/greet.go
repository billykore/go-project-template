package domain

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

// Greet is greet entity.
type Greet struct {
	gorm.Model
	Name string
}

// ErrEmptyGreets indicates no messages was stored in the repository.
var ErrEmptyGreets = errors.New("empty messages")

// GreetRepository defines the methods to interacting with persistence storage used by greet domain.
type GreetRepository interface {
	// GetAll gets all messages.
	GetAll(ctx context.Context) ([]Greet, error)

	// Save saves message.
	Save(ctx context.Context, message Greet) error
}
