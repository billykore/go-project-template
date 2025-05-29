package example

import (
	"github.com/billykore/go-service-tmpl/internal/domain/example"
	"github.com/billykore/go-service-tmpl/internal/pkg/log"
	"github.com/billykore/go-service-tmpl/internal/pkg/pkgerror"
)

type Service struct {
	log  *log.Logger
	repo example.Repository
}

func NewService(repo example.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Get(id int) (example.Entity, error) {
	e, err := s.repo.Get(id)
	if err != nil {
		s.log.Usecase("Get").Errorf("Get failed: %v", err)
		return example.Entity{}, pkgerror.NotFound(ErrEntityNotFound).SetMsg("Example not found.")
	}
	return e, nil
}

func (s *Service) Save(e example.Entity) error {
	err := s.repo.Save(e)
	if err != nil {
		s.log.Usecase("Save").Errorf("Save filed: %v", err)
		return pkgerror.InternalServerError(ErrSaveEntityFailed)
	}
	return nil
}
