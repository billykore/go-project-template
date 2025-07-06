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

func NewService(log *log.Logger, repo example.Repository) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}

func (s *Service) GetEntity(id int) (example.Entity, error) {
	e, err := s.repo.Get(id)
	if err != nil {
		s.log.Usecase("GetEntity").Errorf("GetEntity failed: %v", err)
		return example.Entity{}, pkgerror.NotFound().SetMsg("Example not found.")
	}
	return e, nil
}

func (s *Service) SaveEntity(id int) error {
	err := s.repo.Save(example.Entity{ID: id})
	if err != nil {
		s.log.Usecase("SaveEntity").Errorf("SaveEntity failed: %v", err)
		return pkgerror.InternalServerError()
	}
	return nil
}
