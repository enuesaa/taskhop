package lib

import (
	"github.com/enuesaa/taskhop/lib/archivefx"
	"github.com/enuesaa/taskhop/lib/cmdfx"
	"github.com/enuesaa/taskhop/lib/logfx"
	"github.com/enuesaa/taskhop/lib/procfx"
	"github.com/enuesaa/taskhop/lib/taskfx"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"lib",
	archivefx.Module,
	cmdfx.Module,
	taskfx.Module,
	procfx.Module,
	logfx.Module,
	fx.Provide(New),
)
