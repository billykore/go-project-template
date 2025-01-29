package main

import (
	_ "github.com/billykore/go-service-tmpl/cmd/swagger/docs"
	"github.com/billykore/go-service-tmpl/internal/adapter/http/server"
	"github.com/billykore/go-service-tmpl/pkg/config"
	_ "github.com/joho/godotenv/autoload"
)

type app struct {
	ss *server.Server
}

func newApp(ss *server.Server) *app {
	return &app{
		ss: ss,
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
	c := config.Get()
	a := initApp(c)
	a.ss.Serve()
}
