package commander

import (
	"github.com/enuesaa/taskhop/commander/gql"
	"github.com/enuesaa/taskhop/commander/gqlplayground"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		// server
		fx.Annotate(
			NewServer,
			fx.ParamTags(`group:"routes"`),
		),

		// routes
		fx.Annotate(
			gql.New,
			fx.As(new(Route)),
			fx.ResultTags(`group:"routes"`),
		),
		fx.Annotate(
			gqlplayground.New,
			fx.As(new(Route)),
			fx.ResultTags(`group:"routes"`),
		),
	),
)
