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
		a.usecase.Info(ctx, "polling..")

		appctx, err := a.usecase.Connect(ctx)
		if err != nil {
			return err
		}
		a.usecase.Info(appctx, "START")

		if err := a.run(appctx); err != nil {
			a.usecase.InfoE(appctx, err)
		}
	}
}

func (a *Agent) run(ctx context.Context) error {
	for task := range a.usecase.SubscribeTask(ctx) {
		if task.Err != nil {
			if errors.Is(task.Err, gqlclient.ErrConnectionEnd) {
				a.usecase.Info(ctx, "END")
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
