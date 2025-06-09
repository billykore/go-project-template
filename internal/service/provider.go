package service

import (
	"github.com/billykore/go-service-tmpl/internal/service/example"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	example.NewService,
)
