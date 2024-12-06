package greet

import (
	"context"
	"errors"

	"github.com/billykore/go-service-tmpl/pkg/codes"
	"github.com/billykore/go-service-tmpl/pkg/logger"
	"github.com/billykore/go-service-tmpl/pkg/status"
)

const (
	messageGetHistoryFailed = "Get history failed"
	messageEmptyGreet       = "Empty greet"
	messageSayHelloFailed   = "Say hello failed"
)

type Service struct {
	log   *logger.Logger
	repo  Repository
	email Email
}

func NewService(log *logger.Logger, repo Repository) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}

func (s *Service) History(ctx context.Context) ([]*Response, error) {
	messages, err := s.repo.GetAll(ctx)
	if err != nil && errors.Is(err, ErrEmptyGreets) {
		s.log.Usecase("History").Errorf("Empty messages: %v", err)
		return nil, status.Error(codes.NotFound, messageEmptyGreet)
	}
	if err != nil {
		s.log.Usecase("History").Errorf("Failed to get all messages: %v", err)
		return nil, status.Error(codes.NotFound, messageGetHistoryFailed)
	}
	resp := make([]*Response, len(messages))
	for i, message := range messages {
		resp[i] = &Response{
			Name: message.Name,
		}
	}
	s.log.Usecase("History").Info("Get all messages successful")
	return resp, nil
}

func (s *Service) SayHello(ctx context.Context, req HelloRequest) (*HelloResponse, error) {
	err := s.repo.Save(ctx, Greet{Name: req.Name})
	if err != nil {
		s.log.Usecase("SayHello").Errorf("Save message failed: %v", err)
		return nil, status.Error(codes.Internal, messageSayHelloFailed)
	}
	go func() {
		err = s.email.Send(ctx, EmailData{
			Subject: "Greeting",
			Body:    "Successfully say hello " + req.Name,
		})
		if err != nil {
			s.log.Usecase("SayHello").Errorf("Send email failed: %v", err)
		}
	}()
	s.log.Usecase("SayHello").Infof("Say hello to %s", req.Name)
	return &HelloResponse{Message: "Hello " + req.Name}, nil
}
