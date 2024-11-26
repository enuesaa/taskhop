package runner

import (
	"github.com/enuesaa/taskhop/internal"
	"github.com/enuesaa/taskhop/runner/connect"
	"go.uber.org/fx"
)

func New(commanderAddress string) *fx.App {
	connect.Connect(commanderAddress)

	app := fx.New(
		internal.Module,
		fx.NopLogger,
	)

	return app
}
