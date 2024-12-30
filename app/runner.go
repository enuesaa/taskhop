package app

import (
	"github.com/enuesaa/taskhop/app/connector"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
	"go.uber.org/fx"
)

func NewRunner(config *conf.Config, li lib.Lib, shutdowner fx.Shutdowner) Runner {
	runner := Runner{
		config:     config,
		lib:        li,
		conn:       connector.New(config),
		shutdowner: shutdowner,
	}
	return runner
}

type Runner struct {
	config     *conf.Config
	lib        lib.Lib
	conn       connector.Connector
	shutdowner fx.Shutdowner
}

func (a *Runner) Run() error {
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
