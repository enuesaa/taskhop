package logfx

import (
	"github.com/enuesaa/taskhop/internal/logfx/repository"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		New,
		repository.New,
	),
)
