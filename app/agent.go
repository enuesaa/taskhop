package app

import (
	"context"

	"github.com/enuesaa/taskhop/app/gqlclient"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
)

func NewAgent(config *conf.Config, li *lib.Lib) *Agent {
	agent := &Agent{
		config:     config,
		li:         li,
		usecase:    gqlclient.New(config, li),
	}
	return agent
}

type Agent struct {
	config     *conf.Config
	li         *lib.Lib
	usecase    *gqlclient.UseCase
}

func (a *Agent) Run() error {
	ctx := context.Background()
	for {
		a.li.Log.AppInfo(ctx, "polling...")

		if err := a.usecase.Connect(ctx); err != nil {
			return err
		}
		if err := a.run(ctx); err != nil {
			a.li.Log.AppInfo(ctx, "reset: %s", err.Error())
		}
	}
}

func (a *Agent) run(ctx context.Context) error {
	for task := range a.usecase.SubscribeTask(ctx) {
		if task.Err != nil {
			return task.Err
		}
		if task.IsDownload {
			if err := a.usecase.DownloadAssets(ctx); err != nil {
				return err
			}
		}
		if task.IsUpload {
			if err := a.usecase.UploadAssets(ctx); err != nil {
				return err
			}
		}
		if task.IsCmd {
			if err := a.usecase.Cmd(ctx, task); err != nil {
				return err
			}
		}
		if err := a.usecase.Completed(); err != nil {
			return err
		}
	}
	return nil
}
