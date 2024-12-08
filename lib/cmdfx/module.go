package cmdfx

import (
	"github.com/enuesaa/taskhop/lib/cmdfx/repository"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		New,
		repository.New,
	),
)
