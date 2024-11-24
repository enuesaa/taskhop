package main

import (
	"context"
	"log"
	"net/http"

	"github.com/enuesaa/taskhop/gql"
	"github.com/enuesaa/taskhop/internal/cmdexec"
	"github.com/enuesaa/taskhop/internal/fs"
	"github.com/enuesaa/taskhop/internal/logging"
	"github.com/enuesaa/taskhop/internal/usecase"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			logging.New,
			cmdexec.New,
			fs.New,
			gql.New,
			usecase.New,
			NewServer,
		),
		fx.Invoke(func(lc fx.Lifecycle, s *http.Server) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						if err := s.ListenAndServe(); err != nil {
							log.Printf("Error: %s\n", err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return s.Shutdown(ctx)
				},
			})
		}),
	)
	app.Run()
	// ctx := context.Background()
	// ctx = repos.Log.Use(ctx, "a", "b")
	// repos.Log.Info(ctx, "aa")

	// if err := http.ListenAndServe(":3000", router); err != nil {
	// 	log.Fatalf("Error: %s", err.Error())
	// }
}
