package logfx

import "context"

type taskIdKey struct{}

func (i *LogSrv) Use(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, taskIdKey{}, id)
}
