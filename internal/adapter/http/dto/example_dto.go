package dto

import (
	"github.com/billykore/go-service-tmpl/internal/domain/example"
)

// GetRequest represents a request to retrieve data or resources from the system.
type GetRequest struct {
	ID int
}

// GetResponse represents the response received after a get operation.
type GetResponse struct {
	ID int
}

// NewGetResponse creates a new GetResponse instance.
func NewGetResponse(e example.Entity) *GetResponse {
	return &GetResponse{
		ID: e.ID,
	}
}

// SaveRequest represents a request to save data or resources in the system.
type SaveRequest struct {
	ID int
}
