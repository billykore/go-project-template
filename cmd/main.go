package main

import (
	"github.com/billykore/go-service-tmpl/internal/adapter/http/server"
	"github.com/billykore/go-service-tmpl/internal/pkg/config"
)

type app struct {
	srv *server.Server
}

func newApp(srv *server.Server) *app {
	return &app{
		srv: srv,
	}
}

// main swaggo annotation.
//
//	@title			API Specification
//	@version		1.0
//	@description	Greet service API specification.
//	@termsOfService	https://swagger.io/terms/
//	@contact.name	[Your Name]
//	@contact.url	https://www.swagger.io/support
//	@contact.email	[Your Email]
//	@license.name	Apache 2.0
//	@license.url	https://www.apache.org/licenses/LICENSE-2.0.html
//	@host			id.greet.org
//	@schemes		http https
//	@BasePath		/api/v1
func main() {
	c := config.Load()
	a := initApp(c)
	a.srv.Run()
}
