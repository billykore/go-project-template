package server

import (
	"github.com/billykore/go-service-tmpl/internal/adapter/http/handler"
	"github.com/billykore/go-service-tmpl/internal/pkg/config"
	"github.com/billykore/go-service-tmpl/internal/pkg/log"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

// Server represents the main server struct managing configuration, logging, and routing.
type Server struct {
	cfg            *config.Configs
	log            *log.Logger
	router         *echo.Echo
	exampleHandler *handler.ExampleHandler
}

// New returns new Router.
func New(
	cfg *config.Configs,
	log *log.Logger,
	router *echo.Echo,
	exampleHandler *handler.ExampleHandler,
) *Server {
	return &Server{
		cfg:            cfg,
		log:            log,
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
	srv.log.Infof("running on port [:%v]", port)
	if err := srv.router.Start(":" + port); err != nil {
		srv.log.Fatalf("failed to run on port [:%v]", port)
	}
}
