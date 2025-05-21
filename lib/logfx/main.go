package logfx

import "context"

func New(repository IRepository) ILogSrv {
	return &LogSrv{
		repository: repository,
	}
}

type ILogSrv interface {
	AppInfo(ctx context.Context, format string, a ...any)
	Error(ctx context.Context, err error)
	Info(ctx context.Context, format string, a ...any)
	Debug(ctx context.Context, format string, a ...any)
	Faint(ctx context.Context, format string, a ...any)
}

type LogSrv struct {
	repository IRepository
}
