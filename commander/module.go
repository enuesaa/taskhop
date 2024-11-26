package commander

import (
	"github.com/enuesaa/taskhop/commander/gql"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewServer,
			fx.ParamTags(`group:"routes"`),
		),
		fx.Annotate(
			gql.New,
			fx.As(new(Route)),
			fx.ResultTags(`group:"routes"`),
		),
	),
)
