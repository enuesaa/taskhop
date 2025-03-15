package app

import (
	"time"

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
	if err := a.Connect(); err != nil {
		return err
	}
	if err := a.Register(); err != nil {
		return err
	}

	// if err := a.UnArchive(); err != nil {
	// 	return err
	// }
	for {
		time.Sleep(5 * time.Second)
		if err := a.run(); err != nil {
			return err
		}
	}
}
