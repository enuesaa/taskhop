package gqlclient

import "context"

func (u *UseCase) AppInfo(ctx context.Context, format string, a ...any) {
	u.li.Log.AppInfo(ctx, format, a...)
}

func (u *UseCase) AppDebugE(ctx context.Context, err error) {
	//
}
