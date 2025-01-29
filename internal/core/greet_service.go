package core

import (
	"context"

	"github.com/billykore/go-service-tmpl/internal/domain"
	"github.com/billykore/go-service-tmpl/internal/dto"
	"github.com/billykore/go-service-tmpl/pkg/codes"
	"github.com/billykore/go-service-tmpl/pkg/logger"
	"github.com/billykore/go-service-tmpl/pkg/status"
)

const (
	messageGetHistoryFailed = "Get history failed"
	messageSayHelloFailed   = "Say hello failed"
)

type GreetService struct {
	log  *logger.Logger
	repo domain.GreetRepository
}

func NewGreetService(log *logger.Logger, repo domain.GreetRepository) *GreetService {
	return &GreetService{
		log:  log,
		repo: repo,
	}
}

func (s *GreetService) History(ctx context.Context) ([]*dto.Response, error) {
	messages, err := s.repo.GetAll(ctx)
	if err != nil {
		s.log.Usecase("History").Errorf("Failed to get all messages: %v", err)
		return nil, status.Error(codes.NotFound, messageGetHistoryFailed)
	}
	resp := make([]*dto.Response, len(messages))
	for i, message := range messages {
		resp[i] = &dto.Response{
			Name: message.Name,
		}
	}
	s.log.Usecase("History").Info("Get all messages successful")
	return resp, nil
}

func (s *GreetService) SayHello(ctx context.Context, req dto.HelloRequest) (*dto.HelloResponse, error) {
	err := s.repo.Save(ctx, domain.Greet{Name: req.Name})
	if err != nil {
		s.log.Usecase("SayHello").Errorf("Save message failed: %v", err)
		return nil, status.Error(codes.Internal, messageSayHelloFailed)
	}
	s.log.Usecase("SayHello").Infof("Say hello to %s", req.Name)
	return &dto.HelloResponse{Message: "Hello " + req.Name}, nil
}
