package example

import (
	"github.com/billykore/go-service-tmpl/internal/domain/example"
	"github.com/billykore/go-service-tmpl/internal/pkg/log"
	"github.com/billykore/go-service-tmpl/internal/pkg/pkgerror"
)

type Usecase struct {
	log  *log.Logger
	repo example.Repository
}

func NewUsecase(log *log.Logger, repo example.Repository) *Usecase {
	return &Usecase{
		log:  log,
		repo: repo,
	}
}

func (uc *Usecase) GetEntity(id int) (example.Entity, error) {
	e, err := uc.repo.Get(id)
	if err != nil {
		uc.log.Usecase("GetEntity").Errorf("GetEntity failed: %v", err)
		return example.Entity{}, pkgerror.NotFound().SetMsg("Example not found.")
	}
	return e, nil
}

func (uc *Usecase) SaveEntity(id int) error {
	err := uc.repo.Save(example.Entity{ID: id})
	if err != nil {
		uc.log.Usecase("SaveEntity").Errorf("SaveEntity failed: %v", err)
		return pkgerror.InternalServerError()
	}
	return nil
}
