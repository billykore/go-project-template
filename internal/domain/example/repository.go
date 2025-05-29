package example

// Repository defines a contract for data access and persistence operations.
type Repository interface {
	// Get retrieves an entity from the repository by its ID.
	Get(id int) (Entity, error)

	// Save saves an entity to the repository.
	Save(e Entity) error
}
