package logfx

import (
	"context"
)

func New() I {
	return &Impl{}
}

type I interface {
	Use(ctx context.Context, key string, value interface{}) context.Context
	Info(ctx context.Context, format string, a ...any)
}

type Impl struct{}
