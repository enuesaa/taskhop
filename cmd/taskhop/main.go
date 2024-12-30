package main

import (
	"os"

	"github.com/enuesaa/taskhop/app/commander"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	isDebugMode := os.Getenv("TASKHOP_DEBUG") == "true"

	app := fx.New(
		Module,
		conf.Module,
		lib.Module,
		commander.Module,
		fx.WithLogger(func(logger FxLogger) fxevent.Logger {
			logger.Debug = isDebugMode
			return &logger
		}),
		fx.Invoke(func(cli ICli) error {
			return cli.Launch()
		}),
		fx.Invoke(func(commanderApp commander.App) error {
			return commanderApp.Run()
		}),
	)
	app.Run()
}
