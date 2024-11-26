package gqlplayground

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
)

func Handle() http.HandlerFunc {
	return playground.Handler("graphql", "/graphql")
}
