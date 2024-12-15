package runner

import (
	"github.com/enuesaa/taskhop/app/runner/connector"
	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib"
	"go.uber.org/fx"
)

func New(cl cli.ICli, li lib.Lib, shutdowner fx.Shutdowner) App {
	return App{
		cli:        cl,
		lib:        li,
		conn:       connector.New(cl),
		shutdowner: shutdowner,
	}
}

type App struct {
	cli        cli.ICli
	lib        lib.Lib
	conn       connector.Connector
	shutdowner fx.Shutdowner
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
