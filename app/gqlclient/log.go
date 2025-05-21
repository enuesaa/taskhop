package gqlclient

import "context"

func (u *UseCase) AppInfo(ctx context.Context, format string, a ...any) {
	u.li.Log.AppInfo(ctx, format, a...)
}

func (u *UseCase) AppError(ctx context.Context, err error) {
	u.li.Log.Error(ctx, err)
}

func (u *UseCase) AppDebug(ctx context.Context, format string, a... any) {
	u.li.Log.Debug(ctx, format, a...)
}
