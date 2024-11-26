package gql

import (
	"net/http"
	"github.com/99designs/gqlgen/graphql/handler"
)

func New() *Handler {
	schema := NewExecutableSchema(Config{
		Resolvers: &Resolver{},
	})
	handle := handler.NewDefaultServer(schema)

	return &Handler{handle}
}

type Handler struct {
	http.Handler
}

func (h *Handler) Pattern() string {
	return "/graphql"
}
