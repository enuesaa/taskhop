package runner

import (
	"fmt"

	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/runner/client"
	"github.com/enuesaa/taskhop/runner/connector"
	"go.uber.org/fx"
)

func New(workdir string, connect string) *fx.App {
	address := client.Address(connect)

	app := fx.New(
		internal.Module,
		fx.Supply(address),
		fx.Provide(
			client.New,
			connector.New,
		),
		fx.Invoke(func(con *connector.Connector) error {
			return con.Connect()
		}),
		fx.Invoke(func(con *connector.Connector) error {
			task, err := con.Register()
			if err != nil {
				return err
			}
			fmt.Printf("%+v\n", task)

			if err := con.UnArchive(workdir); err != nil {
				return err
			}

			if err := con.Run(task, workdir); err != nil {
				return err
			}

			return nil
		}),
		fx.NopLogger,
	)

	return app
}
