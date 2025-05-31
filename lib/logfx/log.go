package logfx

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

func (i *LogSrv) Faint(ctx context.Context, format string, a ...any) {
	italic := color.New(color.Faint)
	text := italic.Sprintf(format, a...)

	i.repository.Print(text)
}

func (i *LogSrv) Info(ctx context.Context, format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	i.repository.Print(text)
}
