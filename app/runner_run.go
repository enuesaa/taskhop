package app

import (
	"context"

	"github.com/enuesaa/taskhop/app/gql/model"
)

func (a *Runner) run() error {
	taskres, err := a.conn.GetTask(context.Background())
	if err != nil {
		return err
	}
	task := taskres.Task

	if task.Status != model.TaskStatusProceeding {
		a.lib.Log.AppInfo(context.Background(), "waiting")
		return nil
	}
	a.lib.Log.AppInfo(context.Background(), "started")

	for _, command := range task.Cmds {
		input := model.LogInput{
			Type:   model.LogTypeCommand,
			Output: command,
		}
		if _, err := a.conn.Log(context.Background(), input); err != nil {
			return err
		}
		if err := a.lib.Cmd.Exec(&a.conn, command, a.config.Workdir); err != nil {
			return err
		}
	}

	a.lib.Log.AppInfo(context.Background(), "complete!")
	a.lib.Log.AppInfo(context.Background(), "")

	if _, err := a.conn.Completed(context.Background()); err != nil {
		return err
	}
	return nil
}
