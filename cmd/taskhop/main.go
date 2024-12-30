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
	isDebugMode := os.Getenv("TASKHOP_DEBUG") == "true"

	fxapp := fx.New(
		Module,
		conf.Module,
		lib.Module,
		app.Module,
		fx.WithLogger(func(logger FxLogger) fxevent.Logger {
			logger.Debug = isDebugMode
			return &logger
		}),
		fx.Invoke(func(cli ICli) error {
			return cli.Launch()
		}),
		fx.Invoke(func(commander app.Commander) error {
			return commander.Run()
		}),
	)
	fxapp.Run()
}
