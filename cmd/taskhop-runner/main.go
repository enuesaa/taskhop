package main

import (
	// "github.com/enuesaa/taskhop/app/commander"
	// "github.com/enuesaa/taskhop/app/runner"
	"github.com/enuesaa/taskhop/conf"
	// "github.com/enuesaa/taskhop/lib"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	app := fx.New(
		Module,
		conf.Module,
		// lib.Module,
		// commander.Module,
		// runner.Module,
		fx.WithLogger(func(logger FxLogger) fxevent.Logger {
			return &logger
		}),
		fx.Invoke(func(cli ICli) error {
			return cli.Launch()
		}),
		// fx.Invoke(func(cl cli.ICli, runnerApp runner.App) error {
		// 	return runnerApp.Run()
		// }),
	)
	app.Run()
}
