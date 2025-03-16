package app

import (
	"bytes"
	"context"

	"github.com/enuesaa/taskhop/app/gql/model"
)

func (a *Runner) run(ctx context.Context) error {
	for task := range a.conn.SubscribeTask(ctx) {
		if task.Err != nil {
			return task.Err
		}
		if task.Status == model.TaskStatusPrompt {
			a.lib.Log.AppInfo(ctx, "waiting prompt...")
			continue
		}
		if task.Status == model.TaskStatusDownloadAssets {
			a.lib.Log.AppInfo(ctx, "download assets...")
			var buf bytes.Buffer
			if err := a.conn.GetStorageArchive(&buf); err != nil {
				return err
			}
			if err := a.lib.Arv.UnArchive(&buf, a.config.Workdir); err != nil {
				return err
			}
		}
		if task.Status == model.TaskStatusProceeding {
			a.lib.Log.AppInfo(ctx, "started: %s", task.Cmd)
			if err := a.lib.Cmd.Exec(&a.conn, task.Cmd, a.config.Workdir); err != nil {
				return err
			}
		}
		if _, err := a.conn.Gql.Completed(ctx); err != nil {
			return err
		}
	}
	return nil
}
