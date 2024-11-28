package usecase

import "context"

func (u *UseCase) Connect() error {
	if err := u.conn.DialPolling(); err != nil {
		return err
	}
	if _, err := u.conn.GetHealth(context.Background()); err != nil {
		return err
	}
	return nil
}
