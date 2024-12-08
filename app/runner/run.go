package runner

import (
	"context"

	"github.com/enuesaa/taskhop/app/runner/connector"
)

func (a *App) run(task connector.GetTask_Task) error {
	a.lib.Log.Info(context.Background(), "start cmds...")
	if err := a.lib.Cmd.MultiExec(&a.conn, task.Cmds, a.cli.GetWorkdir()); err != nil {
		return err
	}

	a.lib.Log.Info(context.Background(), "completed")
	_, err := a.conn.Completed(context.Background())
	if err != nil {
		return err
	}
	return nil
}
