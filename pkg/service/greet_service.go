package service

import (
	"context"

	"github.com/billykore/go-service-tmpl/pkg/entity"
	"github.com/billykore/go-service-tmpl/pkg/repository"
	"github.com/billykore/go-service-tmpl/pkg/utils/codes"
	"github.com/billykore/go-service-tmpl/pkg/utils/log"
	"github.com/billykore/go-service-tmpl/pkg/utils/pkgerror"
)

// GreetService provides greeting service business logic.
type GreetService struct {
	log  *log.Logger
	repo repository.GreetRepository
}

func NewGreetService(log *log.Logger, repo repository.GreetRepository) *GreetService {
	return &GreetService{
		log:  log,
		repo: repo,
	}
}

func (s *GreetService) SayHello(ctx context.Context, name string) (entity.Greet, error) {
	user, ok := entity.UserFromContext(ctx)
	if !ok {
		s.log.Usecase("SayHello").Error(ErrGetUserFromContext)
		return entity.Greet{}, pkgerror.New(codes.Unauthenticated, ErrGetUserFromContext).
			SetMsg("Please login to continue.")
	}
	if user.ID == entity.ForbiddenID {
		s.log.Usecase("SayHello").Errorf("Forbidden ID")
		return entity.Greet{}, pkgerror.New(codes.Forbidden, ErrSayHelloFailed).
			SetMsg("User is forbidden to access this resource.")
	}
	if name == entity.DefaultName {
		s.log.Usecase("SayHello").Errorf("Invalid name: %s", name)
		return entity.Greet{}, pkgerror.New(codes.BadRequest, ErrSayHelloFailed).
			SetMsg("Please provide different name.")
	}

	err := s.repo.Save(ctx, entity.Greet{Name: name})
	if err != nil {
		s.log.Usecase("SayHello").Errorf("Save message failed: %v", err)
		return entity.Greet{}, pkgerror.New(codes.Internal, ErrSayHelloFailed)
	}

	s.log.Usecase("SayHello").Infof("Say hello to %s", name)
	return entity.Greet{}, nil
}
