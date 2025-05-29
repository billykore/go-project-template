package example

import "errors"

var (
	// ErrEntityNotFound indicates that the requested entity could not be located.
	ErrEntityNotFound = errors.New("entity not found")

	// ErrSaveEntityFailed indicates that the entity could not be saved.
	ErrSaveEntityFailed = errors.New("failed to save entity")
)
