package logfx

import (
	"context"
	"fmt"
	"log/slog"
)

func (i *LogSrv) Info(ctx context.Context, format string, a ...any) {
	message := fmt.Sprintf(format, a...)

	taskId, ok := ctx.Value(taskIdKey{}).(string)
	if ok {
		message = fmt.Sprintf("[%s] %s", taskId, message)
	}
	slog.Info(message)
}
