package gql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/enuesaa/taskhop/lib"
)

func Handle(c lib.Lib) http.HandlerFunc {
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
