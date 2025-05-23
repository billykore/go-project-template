package internal

import (
	"github.com/billykore/go-service-tmpl/internal/http/handler"
	"github.com/billykore/go-service-tmpl/internal/http/router"
	"github.com/billykore/go-service-tmpl/internal/storage/repo"
	"github.com/billykore/go-service-tmpl/pkg/repository"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repo.NewGreetRepo, wire.Bind(new(repository.GreetRepository), new(*repo.GreetRepo)),
	handler.NewGreetHandler,
	router.New,
)
