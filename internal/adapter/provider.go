package adapter

import (
	"github.com/billykore/go-service-tmpl/internal/adapter/http/handler"
	"github.com/billykore/go-service-tmpl/internal/adapter/http/server"
	"github.com/billykore/go-service-tmpl/internal/adapter/storage/repo"
	"github.com/billykore/go-service-tmpl/internal/domain/example"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repo.NewExampleRepo, wire.Bind(new(example.Repository), new(*repo.ExampleRepo)),
	handler.NewExampleHandler,
	server.New,
)
