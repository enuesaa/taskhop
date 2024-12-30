package main

import "go.uber.org/fx"

var Module = fx.Module(
	"taskhop-runner",
	fx.Provide(
		New,
		NewFxLogger,
	),
)
