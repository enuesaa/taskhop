package runfx

import (
	"github.com/enuesaa/taskhop/internal/runfx/repository"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		New,
		repository.New,
	),
)
