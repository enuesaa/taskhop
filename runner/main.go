package runner

import (
	"fmt"
	"log"

	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/runner/usecase"
	"github.com/enuesaa/taskhop/runner/connector"
	"go.uber.org/fx"
)

func New(workdir string, address string) *fx.App {
	client, err := connector.New(address)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	app := fx.New(
		internal.Module,
		fx.Supply(client),
		fx.Provide(usecase.New),
		fx.Invoke(func(u *usecase.UseCase) error {
			return u.Connect()
		}),
		fx.Invoke(func(u *usecase.UseCase) error {
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
