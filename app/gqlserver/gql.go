package gqlserver

import (
	"github.com/enuesaa/taskhop/app/gqlserver/mutation"
	"github.com/enuesaa/taskhop/app/gqlserver/query"
	"github.com/enuesaa/taskhop/lib"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func (h *Handler) HandleGQL() http.HandlerFunc {
	schema := NewExecutableSchema(Config{
		Resolvers: &Resolver{h.li},
	})

	handle := handler.New(schema)
	handle.AddTransport(transport.Options{})
	handle.AddTransport(transport.GET{})
	handle.AddTransport(transport.POST{})
	handle.AddTransport(transport.MultipartForm{})

	return handle.ServeHTTP
}

type Resolver struct {
	li *lib.Lib
}

func (r *Resolver) Query() QueryResolver {
	return query.New(r.li)
}

func (r *Resolver) Mutation() MutationResolver {
	return mutation.New(r.li)
}
