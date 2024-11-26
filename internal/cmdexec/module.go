package cmdexec

import (
	"github.com/enuesaa/taskhop/internal/cmdexec/repository"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		New,
		repository.New,
	),
)
