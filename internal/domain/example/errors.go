package example

import "errors"

var (
	// ErrNotFound indicates that the requested entity could not be located.
	ErrNotFound = errors.New("entity not found")

	// ErrSaveFailed indicates that the entity could not be saved.
	ErrSaveFailed = errors.New("failed to save entity")
)
