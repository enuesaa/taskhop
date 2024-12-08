package runner

import "context"

func (a *App) Connect() error {
	a.lib.Log.Info(context.Background(), "start polling")
	if err := a.conn.DialPolling(); err != nil {
		return err
	}

	a.lib.Log.Info(context.Background(), "check health")
	if _, err := a.conn.GetHealth(context.Background()); err != nil {
		return err
	}
	return nil
}
