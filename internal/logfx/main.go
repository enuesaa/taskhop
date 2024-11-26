package logfx

import (
	"context"
	"log/slog"
	"os"
)

func init() {
	// TODO remove
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func New() I {
	return &Impl{}
}

type I interface {
	Use(ctx context.Context, key string, value interface{}) context.Context
	Info(ctx context.Context, format string, a ...any)
}

type Impl struct{}
