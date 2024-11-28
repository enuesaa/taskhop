package usecase

import (
	"context"

	"github.com/enuesaa/taskhop/runner/connector"
)

func (u *UseCase) Run(task connector.GetTask_Task) error {
	u.Logi.Info(context.Background(), "start cmds...")
	if err := u.Cmdi.MultiExec(&u.conn, task.Cmds, u.config.Workdir); err != nil {
		return err
	}

	u.Logi.Info(context.Background(), "completed")
	_, err := u.conn.Completed(context.Background())
	if err != nil {
		return err
	}
	return nil
}
