package runner

import (
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
		fx.Invoke(InvokeConnect),
		fx.Invoke(InvokeMain),
		fx.NopLogger,
	)

	return app
}
