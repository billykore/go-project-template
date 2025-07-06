package repo

import (
	"github.com/billykore/go-service-tmpl/internal/domain/example"
)

type ExampleRepo struct {
	// Add your database connection here, e.g. *sql.DB, *gorm.DB, etc.
}

func NewExampleRepo() *ExampleRepo {
	return &ExampleRepo{}
}

func (repo ExampleRepo) Get(id int) (example.Entity, error) {
	if id <= 0 {
		return example.Entity{}, example.ErrNotFound
	}
	return example.Entity{ID: id}, nil
}

func (repo ExampleRepo) Save(e example.Entity) error {
	if e.ID <= 0 {
		return example.ErrSaveFailed
	}
	return nil
}
