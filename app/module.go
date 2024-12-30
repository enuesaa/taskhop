package app

import "go.uber.org/fx"

var Module = fx.Module(
	"app",
	fx.Provide(
		NewCommander,
		NewRunner,
	),
)
