//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/go-service-tmpl/internal"
	"github.com/billykore/go-service-tmpl/pkg/service"
	"github.com/billykore/go-service-tmpl/pkg/utils"
	"github.com/billykore/go-service-tmpl/pkg/utils/config"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func initApp(cfg *config.Configs) *app {
	wire.Build(
		internal.ProviderSet,
		service.ProviderSet,
		utils.ProviderSet,
		echo.New,
		newApp,
	)
	return &app{}
}
