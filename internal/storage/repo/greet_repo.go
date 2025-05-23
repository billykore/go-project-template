package repo

import (
	"context"

	"github.com/billykore/go-service-tmpl/internal/storage/model"
	"github.com/billykore/go-service-tmpl/pkg/entity"
	"github.com/billykore/go-service-tmpl/pkg/service"
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

func (r *GreetRepo) GetAll(ctx context.Context) ([]entity.Greet, error) {
	var greets []entity.Greet
	res := r.db.WithContext(ctx).Find(&greets)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(greets) == 0 {
		return nil, service.ErrEmptyGreets
	}
	return greets, nil
}

func (r *GreetRepo) Save(ctx context.Context, message entity.Greet) error {
	res := r.db.WithContext(ctx).Save(&model.Greet{
		Name: message.Name,
	})
	return res.Error
}
