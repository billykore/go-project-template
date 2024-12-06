package server

import (
	"net/http"

	"github.com/billykore/go-service-tmpl/infra/http/handler"
	"github.com/billykore/go-service-tmpl/pkg/config"
	"github.com/billykore/go-service-tmpl/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Server to run.
type Server struct {
	router *Router
}

// New creates new Server.
func New(router *Router) *Server {
	return &Server{
		router: router,
	}
}

// Serve start the Server.
func (s *Server) Serve() {
	s.router.Run()
}

// Router get all request to handlers and returns the response produce by handlers.
type Router struct {
	cfg          *config.Config
	log          *logger.Logger
	router       *echo.Echo
	greetHandler *handler.GreetHandler
}

// NewRouter returns new Router.
func NewRouter(
	cfg *config.Config,
	log *logger.Logger,
	router *echo.Echo,
	greetHandler *handler.GreetHandler,
) *Router {
	return &Router{
		cfg:          cfg,
		log:          log,
		router:       router,
		greetHandler: greetHandler,
	}
}

func (r *Router) useMiddlewares() {
	r.router.Use(middleware.Logger())
	r.router.Use(middleware.Recover())
}

func (r *Router) swagger() {
	r.router.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (r *Router) run() {
	port := r.cfg.HTTP.Port
	r.log.Infof("running on port ::[:%v]", port)
	if err := r.router.Start(":" + port); err != nil {
		r.log.Fatalf("failed to run on port [::%v]", port)
	}
}

// Run runs the server.
func (r *Router) Run() {
	r.useMiddlewares()
	r.swagger()
	r.setGreetRoutes()
	r.setNeedAuthRoutes()
	r.run()
}

func (r *Router) setGreetRoutes() {
	r.router.GET("/hello", r.greetHandler.History)
	r.router.POST("/hello", r.greetHandler.SayHello)
}

func (r *Router) setNeedAuthRoutes() {
	r.router.GET("/need-auth", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	}, authMiddleware())
}
