package taskfx

import "go.uber.org/fx"

var Module = fx.Module(
	"taskfx",
	fx.Provide(
		New,
		NewRepository,
	),
)
