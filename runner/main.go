package runner

import (
	"fmt"
	"log"

	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/internal/cli"
	"github.com/enuesaa/taskhop/runner/connector"
	"github.com/enuesaa/taskhop/runner/usecase"
	"go.uber.org/fx"
)

func New(config cli.Config) *fx.App {
	client, err := connector.New(config.Address)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	app := fx.New(
		internal.Module,
		fx.Supply(client),
		fx.Provide(usecase.New),
		fx.Invoke(func (u *usecase.UseCase) error {
			return u.Connect()
		}),
		fx.Invoke(func (u *usecase.UseCase) error {
			task, err := u.Register()
			if err != nil {
				return err
			}
			fmt.Printf("%+v\n", task)
		
			if err := u.UnArchive(); err != nil {
				return err
			}
		
			if err := u.Run(task); err != nil {
				return err
			}
		
			return nil
		}),
		fx.NopLogger,
	)

	return app
}
