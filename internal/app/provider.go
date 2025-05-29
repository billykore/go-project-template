package app

import (
	"github.com/billykore/go-service-tmpl/internal/app/example"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	example.NewService,
)
