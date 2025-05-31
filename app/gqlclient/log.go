package gqlclient

import (
	"context"
	"fmt"
)

func (u *UseCase) AppInfo(ctx context.Context, format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	u.info(ctx, text)
}

func (u *UseCase) AppDebug(ctx context.Context, format string, a... any) {
	text := fmt.Sprintf(format, a...)
	u.info(ctx, text)
}

func (u *UseCase) info(ctx context.Context, text string) {
	sessionID, is := u.getSessionID(ctx)
	if !is {
		u.li.Log.Info(ctx, text)
		return
	}
	u.li.Log.Info(ctx, "[%s] %s", sessionID, text)
}
