//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/go-service-tmpl/internal/adapter/http"
	"github.com/billykore/go-service-tmpl/internal/adapter/repository"
	"github.com/billykore/go-service-tmpl/internal/core"
	"github.com/billykore/go-service-tmpl/pkg"
	"github.com/billykore/go-service-tmpl/pkg/config"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func initApp(cfg *config.Config) *app {
	wire.Build(
		http.ProviderSet,
		repository.ProviderSet,
		core.ProviderSet,
		pkg.ProviderSet,
		echo.New,
		newApp,
	)
	return &app{}
}
