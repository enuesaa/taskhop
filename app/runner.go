package app

import (
	"context"
	"bytes"

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
			a.lib.Log.AppInfo(context.Background(), "reset: %s", err.Error())
		}
	}
}

func (a *Runner) UnArchive() error {
	var buf bytes.Buffer
	if err := a.conn.GetStorageArchive(&buf); err != nil {
		return err
	}
	return a.lib.Arv.UnArchive(&buf, a.config.Workdir)
}
