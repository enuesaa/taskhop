package gqlplayground

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
)

func New() *Handler {
	handle := playground.Handler("graphql", "/graphql")

	return &Handler{handle}
}

type Handler struct {
	http.Handler
}

func (h *Handler) Pattern() string {
	return "/graphql/playground"
}
