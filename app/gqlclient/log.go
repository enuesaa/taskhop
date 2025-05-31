package gqlclient

import (
	"context"
	"fmt"
)

func (u *UseCase) Info(ctx context.Context, format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	u.log(ctx, text)
}

func (u *UseCase) InfoE(ctx context.Context, err error) {
	text := fmt.Sprintf("ERROR: %s", err.Error())
	u.log(ctx, text)
}

func (u *UseCase) log(ctx context.Context, text string) {
	sessionID, is := u.getSessionID(ctx)
	if !is {
		u.li.Log.Info(ctx, text)
		return
	}
	u.li.Log.Info(ctx, "[%s] %s", sessionID, text)
}
