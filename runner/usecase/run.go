package usecase

import "github.com/enuesaa/taskhop/runner/connector"

func (u *UseCase) Run(task connector.GetTask_Task) error {
	if err := u.Cmdi.MultiExec(&u.client, task.Cmds, u.Workdir); err != nil {
		return err
	}

	return nil
}