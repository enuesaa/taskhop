package archivefx

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"archivefx",
	fx.Provide(
		New,
		NewRepository,
	),
)
