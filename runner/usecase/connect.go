package usecase

import "context"

func (u *UseCase) Connect() error {
	if err := u.client.DialPolling(); err != nil {
		return err
	}
	if _, err := u.client.GetHealth(context.Background()); err != nil {
		return err
	}
	return nil
}
