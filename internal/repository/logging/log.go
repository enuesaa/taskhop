package logging

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

type logAttrKey struct{}

type LogRepositoryInterface interface {
	Use(ctx context.Context, key string, value interface{}) context.Context
	Info(ctx context.Context, format string, a ...any)
}

type LogRepository struct {}

func (repo *LogRepository) Use(ctx context.Context, key string, value interface{}) context.Context {
	logattr, ok := ctx.Value(logAttrKey{}).(map[string]interface{})
	if !ok {
		logattr = map[string]interface{}{}
	}
	logattr[key] = value

	return context.WithValue(ctx, logAttrKey{}, logattr)
}

func (repo *LogRepository) Info(ctx context.Context, format string, a ...any) {
	logattr, ok := ctx.Value(logAttrKey{}).(map[string]interface{})
	if !ok {
		logattr = map[string]interface{}{}
	}
	message := fmt.Sprintf(format, a...)
	slog.Info(message, slog.Attr{Key: "attr", Value: slog.AnyValue(logattr)})
}
