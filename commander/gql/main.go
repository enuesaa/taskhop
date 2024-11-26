package gql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/enuesaa/taskhop/internal"
)

func Handle(c internal.Container) http.HandlerFunc {
	schema := NewExecutableSchema(Config{
		Resolvers: &Resolver{c},
	})
	handle := handler.NewDefaultServer(schema)

	return handle.ServeHTTP
}
