package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/billykore/go-service-tmpl/internal/adapter/http/server"
	"github.com/billykore/go-service-tmpl/internal/pkg/config"
	"github.com/billykore/go-service-tmpl/internal/pkg/log"
	zlog "github.com/rs/zerolog/log"
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
	log.Configure(c.App.Env)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// run http server
	go a.srv.Run()

	// wait for termination syscalls and doing cleanup operations after received it
	wait := gracefulShutdown(ctx, 3*time.Second, map[string]operation{
		"http-server": func(ctx context.Context) error {
			return a.srv.Shutdown(ctx)
		},
	})

	<-wait
}

// operation is a function that performs some cleanup operation.
type operation func(ctx context.Context) error

// gracefulShutdown waits for termination syscalls and doing cleanup operations after received it
func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	l := zlog.With().Ctx(ctx).Str("caller", "gracefulShutdown").Logger()

	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s
		l.Info().Msg("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			l.Info().Msgf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})
		defer timeoutFunc.Stop()

		var wg sync.WaitGroup
		// Do the operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)

			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				l.Info().Msgf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					l.Info().Msgf("%s: clean up failed: %v", innerKey, err)
					return
				}

				l.Info().Msgf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()
		close(wait)
	}()

	return wait
}
