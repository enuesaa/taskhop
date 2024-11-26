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

func New() *Handler {
	schema := NewExecutableSchema(Config{
		Resolvers: &Resolver{},
	})
	gqhandle := handler.NewDefaultServer(schema)

	return &Handler{gqhandle}
}

type Handler struct {
	http.Handler
}

func (h *Handler) Pattern() string {
	return "/graphql"
}
