package pkg

import (
	"github.com/billykore/go-service-tmpl/internal/pkg/db/postgres"
	"github.com/billykore/go-service-tmpl/internal/pkg/httpclient"
	"github.com/billykore/go-service-tmpl/internal/pkg/log"
	"github.com/billykore/go-service-tmpl/internal/pkg/validation"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	log.New,
	validation.New,
	postgres.New,
	httpclient.New,
)
