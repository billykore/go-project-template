package example

import (
	"github.com/billykore/go-service-tmpl/internal/domain/example"
	"github.com/billykore/go-service-tmpl/internal/pkg/pkgerror"
	"github.com/rs/zerolog/log"
)

type Usecase struct {
	repo example.Repository
}

func NewUsecase(repo example.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (uc *Usecase) GetEntity(id int) (example.Entity, error) {
	l := log.With().Str("usecase", "GetEntity").Logger()

	e, err := uc.repo.Get(id)
	if err != nil {
		l.Error().Err(err).Msg("GetEntity failed")
		return example.Entity{}, pkgerror.NotFound().SetMsg("Example not found.")
	}
	return e, nil
}

func (uc *Usecase) SaveEntity(id int) error {
	l := log.With().Str("usecase", "SaveEntity").Logger()

	err := uc.repo.Save(example.Entity{ID: id})
	if err != nil {
		l.Error().Err(err).Msg("SaveEntity failed")
		return pkgerror.InternalServerError()
	}

	return nil
}
