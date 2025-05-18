package gqlclient

import (
	"context"
	"errors"
	"time"
)

var ErrConnect = errors.New("failed to connect")

func (u *UseCase) Connect(ctx context.Context) error {
	if err := u.dialPolling(); err != nil {
		return err
	}
	if err := u.adap.GetHealth(); err != nil {
		return err
	}
	if err := u.adap.Register(); err != nil {
		return err
	}
	return nil
}

func (u *UseCase) dialPolling() error {
	for range 120 {
		if u.adap.Dial() {
			return nil
		}
		time.Sleep(5 * time.Second)
	}
	return ErrConnect
}
