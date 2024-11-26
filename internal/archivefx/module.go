package archivefx

import (
	"github.com/enuesaa/taskhop/internal/archivefx/repository"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		New,
		repository.New,
	),
)
