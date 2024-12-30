package main

import "go.uber.org/fx"

var Module = fx.Module(
	"taskhop",
	fx.Provide(
		New,
		NewFxLogger,
	),
)
