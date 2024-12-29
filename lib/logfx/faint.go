package logfx

import (
	"context"
	"log"

	"github.com/fatih/color"
)

func (i *LogSrv) Faint(ctx context.Context, format string, a ...any) {
	italic := color.New(color.Faint)
	message := italic.Sprintf(format, a...)
	log.Print(message)
}
