package usecase

import (
	"context"

	"github.com/enuesaa/taskhop/runner/connector"
)

func (u *UseCase) Run(task connector.GetTask_Task) error {
	if err := u.Cmdi.MultiExec(&u.conn, task.Cmds, u.config.Workdir); err != nil {
		return err
	}

	_, err := u.conn.Completed(context.Background())
	if err != nil {
		return err
	}
	return nil
}
