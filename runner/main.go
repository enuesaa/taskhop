package runner

import (
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
		fx.Invoke(func (u *usecase.UseCase) {
			u.Use(workdir)
		}),
		fx.Invoke(InvokeConnect),
		fx.Invoke(InvokeMain),
		fx.NopLogger,
	)

	return app
}
