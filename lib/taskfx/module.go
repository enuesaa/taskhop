package taskfx

import (
	"github.com/enuesaa/taskhop/internal/taskfx/repository"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		New,
		repository.New,
	),
)
