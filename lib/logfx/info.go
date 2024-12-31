package logfx

import (
	"context"
	"fmt"
)

func (i *LogSrv) Info(ctx context.Context, format string, a ...any) {
	text := fmt.Sprintf(format, a...)

	i.repository.Print(text)
}
