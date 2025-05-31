package gqlclient

import (
	"context"
	"fmt"
)

func (u *UseCase) AppInfo(ctx context.Context, format string, a ...any) {
	sessionID, is := u.getSessionID(ctx)
	if !is {
		u.li.Log.Info(ctx, format, a...)
		return
	}
	text := fmt.Sprintf(format, a...)

	u.li.Log.Info(ctx, "[%s] %s", sessionID, text)
}

func (u *UseCase) AppError(ctx context.Context, err error) {
	u.li.Log.Info(ctx, "error: %s", err.Error())
}

func (u *UseCase) AppDebug(ctx context.Context, format string, a... any) {
	u.li.Log.Info(ctx, format, a...)
}
