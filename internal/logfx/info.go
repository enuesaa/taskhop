package logfx

import (
	"context"
	"fmt"
	"log/slog"
)

func (i *Impl) Info(ctx context.Context, format string, a ...any) {
	logattr, ok := ctx.Value(logAttrKey{}).(map[string]interface{})
	if !ok {
		logattr = map[string]interface{}{}
	}
	message := fmt.Sprintf(format, a...)
	slog.Info(message, slog.Attr{Key: "attr", Value: slog.AnyValue(logattr)})
}
