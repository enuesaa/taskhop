package main

import (
	"log"

	"github.com/enuesaa/taskhop/app/commander"
	"github.com/enuesaa/taskhop/app/runner"
	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		cli.Module,
		lib.Module,
		commander.Module,
		runner.Module,
		fx.Invoke(func (cl cli.ICli) error {
			return cl.Launch()
		}),
		fx.Invoke(func(cl cli.ICli, commanderApp commander.IApp, runnerApp runner.IApp) error {
			if cl.IsCommander() {
				return commanderApp.Run()
			}
			return runnerApp.Run()
		}),
		fx.NopLogger,
	)
	app.Run()

	if app.Err() != nil {
		log.Panicf("Error: %s", app.Err().Error())
	}
}
