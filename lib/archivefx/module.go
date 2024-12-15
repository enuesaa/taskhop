package archivefx

import (
	"github.com/enuesaa/taskhop/lib/archivefx/repository"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"archivefx",
	fx.Provide(
		New,
		repository.New,
	),
)
