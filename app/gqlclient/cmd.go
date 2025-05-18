package gqlclient

import (
	"context"
)

func (u *UseCase) Cmd(ctx context.Context, task Task) error {
	u.li.Log.AppInfo(ctx, "started: %s", task.Cmd)

	return u.li.Cmd.Exec(u.LogWriter, task.Cmd, u.config.Workdir)
}
