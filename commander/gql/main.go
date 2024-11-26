package gql

import (
	"net/http"
	"github.com/99designs/gqlgen/graphql/handler"
)

func Handle() http.HandlerFunc {
	schema := NewExecutableSchema(Config{
		Resolvers: &Resolver{},
	})
	gqhandle := handler.NewDefaultServer(schema)

	return gqhandle.ServeHTTP
}

func New() *handler.Server {
	schema := NewExecutableSchema(Config{
		Resolvers: &Resolver{},
	})
	gqhandle := handler.NewDefaultServer(schema)

	return gqhandle
}
