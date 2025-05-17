package main

import (
	"log"
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
		fx.Invoke(func(agent app.Agent) error {
			if err := agent.Run(); err != nil {
				log.Println(err)
				return err
			}
			return nil
		}),
	)
	fxapp.Run()
}
