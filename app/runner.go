package app

import (
	"context"

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
	ctx := context.Background()

	for {
		a.lib.Log.AppInfo(ctx, "polling...")

		if err := a.conn.Connect(ctx); err != nil {
			return err
		}
		if err := a.run(ctx); err != nil {
			a.lib.Log.AppInfo(ctx, "reset: %s", err.Error())
		}
	}
}
