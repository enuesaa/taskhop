package app

import (
	"github.com/enuesaa/taskhop/app/gql/connector"
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
		if err := a.Register(); err != nil {
			return err
		}
		if err := a.run(); err != nil {
			return err
		}
	}
}
