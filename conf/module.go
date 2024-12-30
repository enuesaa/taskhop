package conf

import "go.uber.org/fx"

var Module = fx.Module(
	"conf",
	fx.Provide(New),
)
