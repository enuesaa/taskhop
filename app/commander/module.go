package commander

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"commander",
	fx.Provide(NewCommander),
)
