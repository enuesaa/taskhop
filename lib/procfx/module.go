package procfx

import "go.uber.org/fx"

var Module = fx.Module(
	"procfx",
	fx.Provide(New),
)
