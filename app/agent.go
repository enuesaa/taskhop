package app

import (
	"context"
	"errors"

	"github.com/enuesaa/taskhop/app/gqlclient"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
)

func NewAgent(config *conf.Config, li *lib.Lib) *Agent {
	agent := &Agent{
		config:  config,
		li:      li,
		usecase: gqlclient.New(config, li),
	}
	return agent
}

type Agent struct {
	config  *conf.Config
	li      *lib.Lib
	usecase *gqlclient.UseCase
}

func (a *Agent) Run() error {
	ctx := context.Background()
	for {
		a.usecase.AppInfo(ctx, "polling..")

		if err := a.usecase.Connect(ctx); err != nil {
			return err
		}
		a.usecase.AppInfo(ctx, "started..")

		if err := a.run(ctx); err != nil {
			a.usecase.AppInfo(ctx, "error: %s", err.Error())
		}
	}
}

func (a *Agent) run(ctx context.Context) error {
	for task := range a.usecase.SubscribeTask(ctx) {
		if task.Err != nil {
			if errors.Is(task.Err, gqlclient.ErrConnectionEnd) {
				a.usecase.AppInfo(ctx, "end")
				return nil
			}
			return task.Err
		}
		if task.IsDownload() {
			if err := a.usecase.DownloadAssets(ctx); err != nil {
				return err
			}
		}
		if task.IsUpload() {
			if err := a.usecase.UploadAssets(ctx); err != nil {
				return err
			}
		}
		if task.IsCmd() {
			if err := a.usecase.Cmd(ctx, task); err != nil {
				return err
			}
		}
		if err := a.usecase.Completed(ctx); err != nil {
			return err
		}
	}
	return nil
}
