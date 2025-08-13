package usecase

import (
	"github.com/billykore/go-service-tmpl/internal/usecase/example"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	example.NewUsecase,
)
