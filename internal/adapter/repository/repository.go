package repository

import (
	"github.com/billykore/go-service-tmpl/internal/domain"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewGreetRepo, wire.Bind(new(domain.GreetRepository), new(*GreetRepo)),
)
