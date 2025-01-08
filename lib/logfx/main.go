package logfx

import "context"

func New(repository IRepository) ILogSrv {
	return &LogSrv{
		repository: repository,
	}
}

type ILogSrv interface {
	AppInfo(ctx context.Context, format string, a ...any)
	Info(ctx context.Context, format string, a ...any)
	Faint(ctx context.Context, format string, a ...any)
}

type LogSrv struct {
	repository IRepository
}
