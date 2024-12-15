package logfx

import (
	"context"
)

func New() ILogSrv {
	return &LogSrv{}
}

type ILogSrv interface {
	Use(ctx context.Context, id string) context.Context
	Info(ctx context.Context, format string, a ...any)
}

type LogSrv struct{}
