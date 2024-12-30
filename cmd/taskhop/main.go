package main

import (
	// "github.com/enuesaa/taskhop/app/commander"
	// "github.com/enuesaa/taskhop/app/runner"
	"os"

	"github.com/enuesaa/taskhop/conf"
	// "github.com/enuesaa/taskhop/lib"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	isDebugMode := os.Getenv("TASKHOP_DEBUG") == "true"

	app := fx.New(
		Module,
		conf.Module,
		// lib.Module,
		// commander.Module,
		// runner.Module,
		fx.WithLogger(func(logger FxLogger) fxevent.Logger {
			logger.Debug = isDebugMode
			return &logger
		}),
		fx.Invoke(func(cli ICli) error {
			return cli.Launch()
		}),
		// fx.Invoke(func(cl cli.ICli, commanderApp commander.App) error {
		// 	return commanderApp.Run()
		// }),
	)
	app.Run()
}
