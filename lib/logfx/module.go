package logfx

import "go.uber.org/fx"

var Module = fx.Module(
	"logfx",
	fx.Provide(
		New,
	),
)
