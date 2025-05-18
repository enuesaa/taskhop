package app

import (
	"bytes"
	"context"

	"github.com/enuesaa/taskhop/app/gqlclient"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
	"go.uber.org/fx"
)

func NewAgent(config *conf.Config, li *lib.Lib, shutdowner fx.Shutdowner) *Agent {
	agent := &Agent{
		config:     config,
		li:         li,
		conn:       gqlclient.New(config),
		shutdowner: shutdowner,
	}
	return agent
}

type Agent struct {
	config     *conf.Config
	li         *lib.Lib
	conn       gqlclient.Connector
	shutdowner fx.Shutdowner
}

func (a *Agent) Run() error {
	ctx := context.Background()
	for {
		a.li.Log.AppInfo(ctx, "polling...")

		if err := a.conn.Connect(ctx); err != nil {
			return err
		}
		if err := a.run(ctx); err != nil {
			a.li.Log.AppInfo(ctx, "reset: %s", err.Error())
		}
	}
}

func (a *Agent) run(ctx context.Context) error {
	for task := range a.conn.SubscribeTask(ctx) {
		if task.Err != nil {
			return task.Err
		}
		if task.IsDownload {
			a.li.Log.AppInfo(ctx, "download assets...")
			var buf bytes.Buffer
			if err := a.conn.DownloadAssets(&buf); err != nil {
				return err
			}
			if err := a.li.Arv.UnArchive(&buf, a.config.Workdir); err != nil {
				return err
			}
			a.conn.Log("success!")
		}
		if task.IsUpload {
			a.li.Log.AppInfo(ctx, "upload assets...")
			archive, err := a.li.Arv.Archive(".")
			if err != nil {
				return err
			}
			if err := a.conn.UploadAssets(archive); err != nil {
				return err
			}
			a.conn.Log("success!")
		}
		if task.IsCmd {
			a.li.Log.AppInfo(ctx, "started: %s", task.Cmd)
			if err := a.li.Cmd.Exec(&a.conn.LogWriter, task.Cmd, a.config.Workdir); err != nil {
				return err
			}
		}
		if err := a.conn.MarkCompleted(ctx); err != nil {
			return err
		}
	}
	return nil
}
