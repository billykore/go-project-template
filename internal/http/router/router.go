package router

import (
	"net/http"

	"github.com/billykore/go-service-tmpl/internal/http/handler"
	"github.com/billykore/go-service-tmpl/internal/http/middleware"
	"github.com/billykore/go-service-tmpl/pkg/utils/config"
	"github.com/billykore/go-service-tmpl/pkg/utils/log"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

// Router gets all requests to handlers and returns the response produce by handlers.
type Router struct {
	cfg          *config.Configs
	log          *log.Logger
	router       *echo.Echo
	greetHandler *handler.GreetHandler
}

// New returns new Router.
func New(
	cfg *config.Configs,
	log *log.Logger,
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

// Run runs the server router.
func (r *Router) Run() {
	r.useMiddlewares()
	r.setGreetRoutes()
	r.setNeedAuthRoutes()
	r.run()
}

func (r *Router) useMiddlewares() {
	r.router.Use(echomiddleware.Logger())
	r.router.Use(echomiddleware.Recover())
}

func (r *Router) run() {
	port := r.cfg.App.Port
	r.log.Infof("running on port [:%v]", port)
	if err := r.router.Start(":" + port); err != nil {
		r.log.Fatalf("failed to run on port [:%v]", port)
	}
}

func (r *Router) setGreetRoutes() {
	r.router.POST("/hello", r.greetHandler.SayHello)
}

func (r *Router) setNeedAuthRoutes() {
	na := r.router.Group("/need-auth")
	na.Use(middleware.AuthenticateUser(r.cfg))

	r.router.GET("", func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusOK)
	})
}
