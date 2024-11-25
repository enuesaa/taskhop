package gql

import (
	"github.com/enuesaa/taskhop/gql/mutation"
	"github.com/enuesaa/taskhop/gql/query"
	"github.com/enuesaa/taskhop/internal/usecase"
)

type Resolver struct {
	Usecase usecase.Usecase
}

func (r *Resolver) Query() QueryResolver {
	resolver := query.QueryResolver{
		Usecase: r.Usecase,
	}
	return &resolver
}

func (r *Resolver) Mutation() MutationResolver {
	resolver := mutation.MutationResolver{}
	return &resolver
}
