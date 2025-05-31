package gqlclient

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

func (u *UseCase) Info(ctx context.Context, format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	u.log(ctx, text)
}

func (u *UseCase) InfoE(ctx context.Context, err error) {
	err = u.maperr(err)
	text := fmt.Sprintf("ERROR: %s", err.Error())
	u.log(ctx, text)
}

func (u *UseCase) maperr(err error) error {
	var urlerr *url.Error
	if errors.As(err, &urlerr) {
		return fmt.Errorf("connection error")
	}
	return err
}

func (u *UseCase) log(ctx context.Context, text string) {
	sessionID, is := u.getSessionID(ctx)
	if !is {
		u.li.Log.Info(ctx, text)
		return
	}
	u.li.Log.Info(ctx, "[%s] %s", sessionID, text)
}
