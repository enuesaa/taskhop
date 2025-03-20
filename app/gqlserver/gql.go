package gqlserver

import (
	"github.com/enuesaa/taskhop/app/gqlserver/mutation"
	"github.com/enuesaa/taskhop/app/gqlserver/query"
	"github.com/enuesaa/taskhop/lib"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func HandleGQL(c lib.Lib) http.HandlerFunc {
	schema := NewExecutableSchema(Config{
		Resolvers: &Resolver{c},
	})

	handle := handler.New(schema)

	handle.AddTransport(transport.Options{})
	handle.AddTransport(transport.GET{})
	handle.AddTransport(transport.POST{})
	handle.AddTransport(transport.MultipartForm{})

	return handle.ServeHTTP
}

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
