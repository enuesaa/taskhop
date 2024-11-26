package runner

import (
	"github.com/enuesaa/taskhop/internal/cmdexec"
	"github.com/enuesaa/taskhop/internal/fs"
	"github.com/enuesaa/taskhop/internal/logging"
	"github.com/enuesaa/taskhop/runner/connect"
	"go.uber.org/fx"
)

func New(commanderAddress string) *fx.App {
	connect.Connect(commanderAddress)

	app := fx.New(
		fx.Provide(
			logging.New,
			cmdexec.New,
			fs.New,
		),
		fx.NopLogger,
	)

	return app
}
