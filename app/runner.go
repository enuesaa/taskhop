package app

import (
	"bytes"
	"context"

	"github.com/enuesaa/taskhop/app/gqlclient"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
	"go.uber.org/fx"
)

func NewRunner(config *conf.Config, li lib.Lib, shutdowner fx.Shutdowner) Runner {
	runner := Runner{
		config:     config,
		lib:        li,
		conn:       gqlclient.New(config),
		shutdowner: shutdowner,
	}
	return runner
}

type Runner struct {
	config     *conf.Config
	lib        lib.Lib
	conn       gqlclient.Connector
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

func (a *Runner) run(ctx context.Context) error {
	for task := range a.conn.SubscribeTask(ctx) {
		if task.Err != nil {
			return task.Err
		}
		if task.IsDownload {
			a.lib.Log.AppInfo(ctx, "download assets...")
			var buf bytes.Buffer
			if err := a.conn.DownloadAssets(&buf); err != nil {
				return err
			}
			if err := a.lib.Arv.UnArchive(&buf, a.config.Workdir); err != nil {
				return err
			}
		}
		if task.IsCmd {
			a.lib.Log.AppInfo(ctx, "started: %s", task.Cmd)
			if err := a.lib.Cmd.Exec(&a.conn, task.Cmd, a.config.Workdir); err != nil {
				return err
			}
		}
		if err := a.conn.MarkCompleted(ctx); err != nil {
			return err
		}
	}
	return nil
}
