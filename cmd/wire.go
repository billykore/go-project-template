//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/go-service-tmpl/internal/adapter/http"
	"github.com/billykore/go-service-tmpl/internal/adapter/storage"
	internalapp "github.com/billykore/go-service-tmpl/internal/app"
	"github.com/billykore/go-service-tmpl/internal/pkg"
	"github.com/billykore/go-service-tmpl/internal/pkg/config"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func initApp(cfg *config.Configs) *app {
	wire.Build(
		http.ProviderSet,
		pkg.ProviderSet,
		internalapp.ProviderSet,
		storage.ProviderSet,
		echo.New,
		newApp,
	)
	return &app{}
}
