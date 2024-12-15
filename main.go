package main

import (
	"github.com/enuesaa/taskhop/app/commander"
	"github.com/enuesaa/taskhop/app/runner"
	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	logger := cli.Logger{
		Debug: false,
	}

	app := fx.New(
		cli.Module,
		lib.Module,
		commander.Module,
		runner.Module,
		fx.WithLogger(func () fxevent.Logger {
			return &logger
		}),
		fx.Invoke(func (cl cli.ICli) error {
			return cl.Launch()
		}),
		fx.Invoke(func (cl cli.ICli) {
			logger.Debug = cl.IsDebug()
		}),
		fx.Invoke(func(cl cli.ICli, commanderApp commander.App, runnerApp runner.App) error {
			if cl.IsCommander() {
				return commanderApp.Run()
			}
			return runnerApp.Run()
		}),
	)
	app.Run()
}
