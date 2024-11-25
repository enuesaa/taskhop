package commander

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/enuesaa/taskhop/gql"
	"github.com/enuesaa/taskhop/internal/cmdexec"
	"github.com/enuesaa/taskhop/internal/cmdsfile"
	"github.com/enuesaa/taskhop/internal/fs"
	"github.com/enuesaa/taskhop/internal/logging"
	"github.com/enuesaa/taskhop/internal/runnermg"
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
		fx.Provide(
			logging.New,
			cmdexec.New,
			fs.New,
			cmdsfile.New,
			gql.New,
			runnermg.New,
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
		fx.NopLogger,
	)

	return app
}
