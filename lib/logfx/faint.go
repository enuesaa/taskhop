package logfx

import (
	"context"

	"github.com/fatih/color"
)

func (i *LogSrv) Faint(ctx context.Context, format string, a ...any) {
	italic := color.New(color.Faint)
	text := italic.Sprintf(format, a...)

	i.repository.Print(text)
}
