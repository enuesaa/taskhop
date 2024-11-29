package logfx

import (
	"context"
)

func New() I {
	return &Impl{}
}

type I interface {
	Use(ctx context.Context, id string) context.Context
	Info(ctx context.Context, format string, a ...any)
}

type Impl struct{}
