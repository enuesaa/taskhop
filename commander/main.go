package commander

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/internal/cmdexec"
	"github.com/enuesaa/taskhop/internal/cmdsfile"
	"github.com/enuesaa/taskhop/internal/logging"
	"github.com/enuesaa/taskhop/internal/runnermg"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func New() *fx.App {
	file, err := cmdsfile.New().Read("testdata/cmds.yml")
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	fmt.Printf("%+v\n", file)

	if err := cmdsfile.New().Validate(file); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	app := fx.New(
		cmdexec.Module,		
		fx.Provide(
			cmdsfile.New,
			logging.New,
			runnermg.New,
			internal.NewContainer,
			NewRouter,
		),
		fx.Invoke(func(lc fx.Lifecycle, router *chi.Mux) {
			server := &http.Server{
				Addr:    ":3000",
				Handler: router,
			}
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						if err := server.ListenAndServe(); err != nil {
							log.Printf("Error: %s\n", err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return server.Shutdown(ctx)
				},
			})
		}),
		// fx.NopLogger,
	)

	return app
}
