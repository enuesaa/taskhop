package app

import (
	"context"
	"fmt"
	"time"

	"github.com/enuesaa/taskhop/app/gql/model"
)

func (a *Runner) run() error {
	ctx := context.Background()

	waitingCount := 0

	for {
		taskres, err := a.conn.GetTask(ctx)
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
			if _, err := a.conn.Completed(ctx); err != nil {
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
			if _, err := a.conn.Completed(ctx); err != nil {
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
