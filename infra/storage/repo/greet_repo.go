package repo

import (
	"context"

	"github.com/billykore/go-service-tmpl/domain/greet"
	"gorm.io/gorm"
)

type GreetRepo struct {
	db *gorm.DB
}

func NewGreetRepo(db *gorm.DB) *GreetRepo {
	return &GreetRepo{
		db: db,
	}
}

func (r *GreetRepo) GetAll(ctx context.Context) ([]greet.Greet, error) {
	var greets []greet.Greet
	res := r.db.WithContext(ctx).Find(&greets)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(greets) == 0 {
		return nil, greet.ErrEmptyGreets
	}
	return greets, nil
}

func (r *GreetRepo) Save(ctx context.Context, message greet.Greet) error {
	res := r.db.WithContext(ctx).Save(&message)
	return res.Error
}
