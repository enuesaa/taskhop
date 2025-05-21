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

func (i *LogSrv) Faint(ctx context.Context, format string, a ...any) {
	italic := color.New(color.Faint)
	text := italic.Sprintf(format, a...)

	i.repository.Print(text)
}

func (i *LogSrv) Info(ctx context.Context, format string, a ...any) {
	text := fmt.Sprintf(format, a...)

	i.repository.Print(text)
}

func (i *LogSrv) Error(ctx context.Context, err error) {
	i.Faint(ctx, "[ERROR] %s", err.Error())
}

func (i *LogSrv) Debug(ctx context.Context, format string, a ...any) {
	i.Faint(ctx, "[DEBUG] "+format, a...)
}
