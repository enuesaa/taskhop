package main

import (
	"os"

	"github.com/enuesaa/taskhop/app"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	fxapp := fx.New(
		fx.Provide(New),
		conf.Module,
		lib.Module,
		app.Module,
		fx.WithLogger(func(logger app.Logger) fxevent.Logger {
			logger.Debug = os.Getenv("TASKHOP_DEBUG") == "true"
			return &logger
		}),
		fx.Invoke(func(cli ICli) error {
			return cli.Launch()
		}),
		fx.Invoke(func(runner app.Runner) error {
			return runner.Run()
		}),
	)
	fxapp.Run()
}
