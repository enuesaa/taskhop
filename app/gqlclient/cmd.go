package gqlclient

import "context"

func (u *UseCase) Cmd(ctx context.Context, task Task) error {
	u.Info(ctx, "cmd: %s", task.Cmd())

	return u.li.Cmd.Exec(u.LogWriter, task.Cmd(), u.config.Workdir)
}
