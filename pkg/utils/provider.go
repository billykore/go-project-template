package utils

import (
	"github.com/billykore/go-service-tmpl/pkg/utils/db/postgres"
	"github.com/billykore/go-service-tmpl/pkg/utils/httpclient"
	"github.com/billykore/go-service-tmpl/pkg/utils/log"
	"github.com/billykore/go-service-tmpl/pkg/utils/validation"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	log.NewLogger,
	validation.New,
	postgres.New,
	httpclient.New,
)
