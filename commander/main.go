package commander

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/internal/cmdexec"
	"github.com/enuesaa/taskhop/internal/cmdsfile"
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
		Module,
		fx.Provide(
			internal.NewContainer,
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
		// fx.NopLogger,
	)

	return app
}
