package runner

import (
	"github.com/enuesaa/taskhop/app/runner/connector"
	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib"
)

type App struct {
	cli cli.ICli
	lib lib.Lib
	conn connector.Connector
}

func (a *App) Run() error {
	for {
		if err := a.Connect(); err != nil {
			return err
		}
		task, err := a.Register()
		if err != nil {
			return err
		}
	
		if err := a.UnArchive(); err != nil {
			return err
		}
		if err := a.run(task); err != nil {
			return err
		}
	}
}
