//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/go-service-tmpl/internal/adapter"
	"github.com/billykore/go-service-tmpl/internal/pkg"
	"github.com/billykore/go-service-tmpl/internal/pkg/config"
	"github.com/billykore/go-service-tmpl/internal/service"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func initApp(cfg *config.Configs) *app {
	wire.Build(
		adapter.ProviderSet,
		pkg.ProviderSet,
		service.ProviderSet,
		echo.New,
		newApp,
	)
	return &app{}
}
