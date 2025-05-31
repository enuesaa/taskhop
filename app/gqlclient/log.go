package gqlclient

import "context"

func (u *UseCase) AppInfo(ctx context.Context, format string, a ...any) {
	u.li.Log.Info(ctx, format, a...)
}

func (u *UseCase) AppError(ctx context.Context, err error) {
	u.li.Log.Info(ctx, "error: %s", err.Error())
}

func (u *UseCase) AppDebug(ctx context.Context, format string, a... any) {
	u.li.Log.Info(ctx, format, a...)
}
