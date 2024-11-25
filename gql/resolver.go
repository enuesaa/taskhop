package gql

import (
	"github.com/enuesaa/taskhop/gql/mutation"
	"github.com/enuesaa/taskhop/gql/query"
)

type Resolver struct {}

func (r *Resolver) Query() QueryResolver {
	resolver := query.QueryResolver{}
	return &resolver
}

func (r *Resolver) Mutation() MutationResolver {
	resolver := mutation.MutationResolver{}
	return &resolver
}
