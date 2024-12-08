package internal

import (
	"github.com/enuesaa/taskhop/internal/archivefx"
	"github.com/enuesaa/taskhop/internal/cmdfx"
	"github.com/enuesaa/taskhop/internal/logfx"
	"github.com/enuesaa/taskhop/internal/taskfx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	archivefx.Module,
	cmdfx.Module,
	taskfx.Module,
	logfx.Module,
	fx.Provide(NewContainer),
)
