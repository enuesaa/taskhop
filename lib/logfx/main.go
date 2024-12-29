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
	Info(ctx context.Context, format string, a ...any)
}

type LogSrv struct{}
