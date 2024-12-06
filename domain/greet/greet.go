package greet

import (
	"context"
	"errors"
)

// ErrEmptyGreets indicates no messages was stored in the repository.
var ErrEmptyGreets = errors.New("empty greets")

// Repository defines the methods to interacting with persistence storage used by greet domain.
type Repository interface {
	// GetAll gets all messages.
	GetAll(ctx context.Context) ([]Greet, error)

	// Save saves message.
	Save(ctx context.Context, message Greet) error
}

// Email defines the behavior for sending email messages.
type Email interface {
	// Send sends an email with the given EmailData.
	// It takes a context to allow for cancellation and timeout handling.
	// Returns an error if the email fails to send.
	Send(ctx context.Context, data EmailData) error
}
