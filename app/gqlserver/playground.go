package gqlserver

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
)

func (h *Handler) HandlePlayground() http.HandlerFunc {
	return playground.Handler("graphql", "/graphql")
}
