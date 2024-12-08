package runner

import (
	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/internal/cli"
	"github.com/enuesaa/taskhop/runner/connector"
	"github.com/enuesaa/taskhop/runner/usecase"
	"go.uber.org/fx"
)

func New(config cli.Config) *fx.App {
	app := fx.New(
		internal.Module,
		fx.Supply(config),
		fx.Provide(connector.New, usecase.New),
		fx.Invoke(func (u usecase.UseCase) error {
			for {
				if err := u.Connect(); err != nil {
					return err
				}
				task, err := u.Register()
				if err != nil {
					return err
				}
			
				if err := u.UnArchive(); err != nil {
					return err
				}
				if err := u.Run(task); err != nil {
					return err
				}
			}
		}),
		fx.NopLogger,
	)

	return app
}
