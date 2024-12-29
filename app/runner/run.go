package runner

import (
	"context"

	"github.com/enuesaa/taskhop/app/commander/gql/model"
	"github.com/enuesaa/taskhop/app/runner/connector"
)

// TODO refactor
func (a *App) run(task connector.GetTask_Task) error {
	a.lib.Log.AppInfo(context.Background(), "started")
	for _, command := range task.Cmds {
		input := model.LogInput{
			Type: model.LogTypeCommand,
			Output: command,
		}
		if _, err := a.conn.Log(context.Background(), input); err != nil {
			return err
		}
		if err := a.lib.Cmd.Exec(&a.conn, command, a.cli.GetWorkdir()); err != nil {
			return err
		}
	}

	a.lib.Log.AppInfo(context.Background(), "complete!")
	a.lib.Log.AppInfo(context.Background(), "")
	_, err := a.conn.Completed(context.Background())
	if err != nil {
		return err
	}
	return nil
}
