package runner

import "context"

func (a *Runner) Connect() error {
	a.lib.Log.AppInfo(context.Background(), "polling...")
	if err := a.conn.DialPolling(); err != nil {
		return err
	}
	if _, err := a.conn.GetHealth(context.Background()); err != nil {
		return err
	}
	return nil
}
