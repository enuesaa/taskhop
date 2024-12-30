package runner

import (
	"github.com/enuesaa/taskhop/app/runner/connector"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
	"go.uber.org/fx"
)

func New(config *conf.Config, li lib.Lib, shutdowner fx.Shutdowner) App {
	return App{
		config:     config,
		lib:        li,
		conn:       connector.New(config),
		shutdowner: shutdowner,
	}
}

type App struct {
	config     *conf.Config
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
