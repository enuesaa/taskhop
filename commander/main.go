package commander

import (
	"context"
	"log"
	"net/http"

	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/internal/archivefx"
	"github.com/enuesaa/taskhop/internal/cli"
	"github.com/enuesaa/taskhop/internal/taskfx"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func New(config cli.Config) *fx.App {
	app := fx.New(
		internal.Module,
		fx.Provide(NewRouter),
		fx.Invoke(func(taski taskfx.I) error {
			taski.Use(config.Workdir)
			task, err := taski.Read()
			if err != nil {
				log.Printf("Error: %s", err.Error())
				return err
			}
			if err := taski.Validate(task); err != nil {
				log.Printf("Error: %s", err.Error())
				return err
			}
			return nil
		}),
		fx.Invoke(func(arvi archivefx.I) error {
			arvi.Use(config.Workdir)
			return nil
		}),
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
		fx.NopLogger,
	)

	return app
}
