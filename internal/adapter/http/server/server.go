package server

import (
	"context"

	"github.com/billykore/go-service-tmpl/internal/adapter/http/handler"
	"github.com/billykore/go-service-tmpl/internal/pkg/config"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

// Server represents the main server struct managing configuration, logging, and routing.
type Server struct {
	cfg            *config.Configs
	router         *echo.Echo
	exampleHandler *handler.ExampleHandler
}

// New returns new Router.
func New(
	cfg *config.Configs,
	router *echo.Echo,
	exampleHandler *handler.ExampleHandler,
) *Server {
	return &Server{
		cfg:            cfg,
		router:         router,
		exampleHandler: exampleHandler,
	}
}

// Run runs the server router.
func (srv *Server) Run() {
	srv.useMiddlewares()
	srv.registerRoutes()
	srv.run()
}

func (srv *Server) useMiddlewares() {
	srv.router.Use(echomiddleware.Logger())
	srv.router.Use(echomiddleware.Recover())
}

func (srv *Server) run() {
	port := srv.cfg.App.Port
	if err := srv.router.Start(":" + port); err != nil {
		log.Panic().Err(err).Msg("Failed to start server")
	}
}

func (hs *Server) Shutdown(ctx context.Context) error {
	return hs.router.Shutdown(ctx)
}
