package commander

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/internal/cli"
	"github.com/enuesaa/taskhop/internal/runfx"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func New(config cli.Config) *fx.App {
	app := fx.New(
		internal.Module,
		fx.Supply(config),
		fx.Provide(NewRouter),
		fx.Invoke(func(c internal.Container) error {
			task, err := c.Taski.Read()
			if err != nil {
				log.Printf("Error: %s", err.Error())
				return err
			}
			if err := c.Taski.Validate(task); err != nil {
				log.Printf("Error: %s", err.Error())
				return err
			}
			c.Logi.Info(context.Background(), "started")
			return nil
		}),
		fx.Invoke(func(c internal.Container, shutdowner fx.Shutdowner) error {
			go func() {
				for status := range c.Runi.Subscribe() {
					fmt.Println("ddd")
					if status == runfx.StatusCompleted {
						shutdowner.Shutdown()
						break
					}
				}
			}()
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
