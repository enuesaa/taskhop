package gqlserver

import (
	"github.com/enuesaa/taskhop/app/gqlserver/mutation"
	"github.com/enuesaa/taskhop/app/gqlserver/query"
	"github.com/enuesaa/taskhop/lib"
)

type Resolver struct {
	lib.Lib
}

func (r *Resolver) Query() QueryResolver {
	resolver := query.QueryResolver{
		Lib: r.Lib,
	}
	return &resolver
}

func (r *Resolver) Mutation() MutationResolver {
	resolver := mutation.MutationResolver{
		Lib: r.Lib,
	}
	return &resolver
}
