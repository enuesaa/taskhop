package internal

import (
	"github.com/enuesaa/taskhop/internal/cmdfx"
	"github.com/enuesaa/taskhop/internal/logfx"
	"github.com/enuesaa/taskhop/internal/runfx"
	"github.com/enuesaa/taskhop/internal/taskfx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	cmdfx.Module,
	taskfx.Module,
	runfx.Module,
	logfx.Module,
	fx.Provide(NewContainer),
)
