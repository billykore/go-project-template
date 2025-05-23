package model

import "gorm.io/gorm"

// Greet represents a model for managing greeting-related data in the database.
type Greet struct {
	gorm.Model
	Name string
}
