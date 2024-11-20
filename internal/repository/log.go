package repository

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

type LogRepositoryInterface interface {
	Use(ctx context.Context, key string, value interface{}) context.Context
	// otel 対応をするのであれば ctx をここで渡し span id を取り出すのが良さそう
	Info(ctx context.Context, format string, a ...any)
}

type LogRepository struct {}

func (repo *LogRepository) Use(ctx context.Context, key string, value interface{}) context.Context {
	logattr, ok := ctx.Value("logattr").(map[string]interface{})
	if !ok {
		logattr = map[string]interface{}{}
	}
	logattr[key] = value

	return context.WithValue(ctx, "logattr", logattr)
}

func (repo *LogRepository) Info(ctx context.Context, format string, a ...any) {
	logattr, ok := ctx.Value("logattr").(map[string]interface{})
	if !ok {
		logattr = map[string]interface{}{}
	}
	message := fmt.Sprintf(format, a...)
	slog.Info(message, slog.Attr{Key: "attr", Value: slog.AnyValue(logattr)})
}
