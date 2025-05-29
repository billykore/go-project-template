package http

import (
	"github.com/billykore/go-service-tmpl/internal/adapter/http/handler"
	"github.com/billykore/go-service-tmpl/internal/adapter/http/server"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	handler.NewExampleHandler,
	server.New,
)
