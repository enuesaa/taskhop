package runner

import (
	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/runner/connector"
	"go.uber.org/fx"
)

func New(commanderAddress string) *fx.App {
	app := fx.New(
		internal.Module,
		fx.Provide(connector.New),
		fx.Invoke(func(con *connector.Connector) error {
			return con.Connect(commanderAddress)
		}),
		// fx.NopLogger,
	)

	return app
}
