package app

import (
	"context"
	"time"

	"github.com/enuesaa/taskhop/app/gql/model"
)

func (a *Runner) run() error {
	ctx := context.Background()

	for {
		taskres, err := a.conn.GetTask(ctx)
		if err != nil {
			return err
		}
		task := taskres.Task

		if task.Status == model.TaskStatusPrompt {
			a.lib.Log.AppInfo(ctx, "waiting prompt...")
			time.Sleep(5 * time.Second)
			continue
		}
		if task.Status == model.TaskStatusCompleted {
			time.Sleep(5 * time.Second)
			continue
		}

		switch task.Status {
		case model.TaskStatusDownloadAssets:
			if err := a.UnArchive(); err != nil {
				return err
			}
		case model.TaskStatusProceeding:
			for _, command := range task.Cmds {
				a.lib.Log.AppInfo(ctx, "started: %s", command)
				input := model.LogInput{
					Type:   model.LogTypeCommand,
					Output: command,
				}
				if _, err := a.conn.Log(ctx, input); err != nil {
					return err
				}
				if err := a.lib.Cmd.Exec(&a.conn, command, a.config.Workdir); err != nil {
					return err
				}
				a.lib.Log.AppInfo(ctx, "completed: %s", command)
			}
		}
		if _, err := a.conn.Completed(ctx); err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}
}
