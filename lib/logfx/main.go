package logfx

import (
	"context"
	"log"
)

func init() {
	log.SetFlags(0)
}

func New() ILogSrv {
	return &LogSrv{}
}

type ILogSrv interface {
	Use(ctx context.Context, id string) context.Context
	AppInfo(ctx context.Context, format string, a ...any)
	Info(ctx context.Context, format string, a ...any)
	Faint(ctx context.Context, format string, a ...any)
}

type LogSrv struct{}
