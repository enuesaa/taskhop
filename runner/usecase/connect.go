package usecase

import "context"

func (u *UseCase) Connect() error {
	u.Logi.Info(context.Background(), "start polling")
	if err := u.conn.DialPolling(); err != nil {
		return err
	}

	u.Logi.Info(context.Background(), "check health")
	if _, err := u.conn.GetHealth(context.Background()); err != nil {
		return err
	}
	return nil
}
