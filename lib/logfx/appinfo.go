package logfx

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

func (i *LogSrv) AppInfo(ctx context.Context, format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	message := color.YellowString(text)

	i.repository.Print(message)
}
