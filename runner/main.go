package runner

import (
	"fmt"

	"github.com/enuesaa/taskhop/runner/client"
	"github.com/enuesaa/taskhop/runner/connector"
	"go.uber.org/fx"
)

func New(commanderAddress string) *fx.App {
	address := client.Address(commanderAddress)

	app := fx.New(
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
			return nil
		}),
		fx.NopLogger,
	)

	return app
}
