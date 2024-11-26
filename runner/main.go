package runner

import (
	"fmt"

	"github.com/enuesaa/taskhop/runner/connector"
	"go.uber.org/fx"
)

func New(commanderAddress string) *fx.App {
	app := fx.New(
		fx.Provide(connector.New),
		fx.Invoke(func(con *connector.Connector) error {
			return con.Connect(commanderAddress)
		}),
		fx.Invoke(func(con *connector.Connector) error {
			task, err := con.Receive(commanderAddress)
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
