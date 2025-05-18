package gqlclient

import "context"

func (u *UseCase) AppInfo(ctx context.Context, format string, a ...any) {
	u.li.Log.AppInfo(ctx, format, a...)
}
