package app

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/enuesaa/taskhop/app/gql/model"
)

func (a *Runner) run(ctx context.Context) error {
	waitingCount := 0

	for {
		taskres, err := a.conn.Gql.GetTask(ctx)
		if err != nil {
			return err
		}
		task := taskres.Task

		switch task.Status {
		case model.TaskStatusPrompt:
			a.lib.Log.AppInfo(ctx, "waiting prompt...")
			waitingCount++
		case model.TaskStatusDownloadAssets:
			a.lib.Log.AppInfo(ctx, "download assets...")
			if err := a.UnArchive(); err != nil {
				return err
			}
			if _, err := a.conn.Gql.Completed(ctx); err != nil {
				return err
			}
		case model.TaskStatusProceeding:
			a.lib.Log.AppInfo(ctx, "started: %s", task.Cmd)
			if err := a.lib.Cmd.Exec(&a.conn, task.Cmd, a.config.Workdir); err != nil {
				return err
			}
			a.lib.Log.AppInfo(ctx, "completed: %s", task.Cmd)
			if _, err := a.conn.Gql.Completed(ctx); err != nil {
				return err
			}
			waitingCount = 0
		}

		time.Sleep(2 * time.Second)
		if waitingCount > 100 {
			return fmt.Errorf("reset")
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
