package runner

import (
	"github.com/enuesaa/taskhop/gql"
	"github.com/enuesaa/taskhop/internal/cmdexec"
	"github.com/enuesaa/taskhop/internal/fs"
	"github.com/enuesaa/taskhop/internal/logging"
	"github.com/enuesaa/taskhop/internal/usecase"
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
			gql.New,
			usecase.New,
		),
		fx.NopLogger,
	)

	return app
}
