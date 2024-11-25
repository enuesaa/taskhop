package runner

import (
	"github.com/enuesaa/taskhop/gql"
	"github.com/enuesaa/taskhop/internal/cmdexec"
	"github.com/enuesaa/taskhop/internal/fs"
	"github.com/enuesaa/taskhop/internal/logging"
	"github.com/enuesaa/taskhop/internal/usecase"
	"go.uber.org/fx"
)

func New() *fx.App {
	Polling()
	fetch()

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
