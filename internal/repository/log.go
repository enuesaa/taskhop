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
	// otel 対応をするのであれば ctx をここで渡し span id を取り出すのが良さそう
	Info(ctx context.Context, format string, a ...any)
}

type LogRepository struct {}

func (repo *LogRepository) Info(ctx context.Context, format string, a ...any) {
	message := fmt.Sprintf(format, a...)
	slog.Info(message)
}
