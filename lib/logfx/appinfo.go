package logfx

import (
	"context"
	"log"

	"github.com/fatih/color"
)

func (i *LogSrv) AppInfo(ctx context.Context, format string, a ...any) {
	message := color.YellowString("*** "+format, a...)
	log.Print(message)
}
