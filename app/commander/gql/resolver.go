package gql

import (
	"github.com/enuesaa/taskhop/commander/gql/mutation"
	"github.com/enuesaa/taskhop/commander/gql/query"
	"github.com/enuesaa/taskhop/internal"
)

type Resolver struct {
	internal.Container
}

func (r *Resolver) Query() QueryResolver {
	resolver := query.QueryResolver{
		Container: r.Container,
	}
	return &resolver
}

func (r *Resolver) Mutation() MutationResolver {
	resolver := mutation.MutationResolver{
		Container: r.Container,
	}
	return &resolver
}
