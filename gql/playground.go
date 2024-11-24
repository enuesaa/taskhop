package gql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
)

func HandlePlayground() http.HandlerFunc {
	return playground.Handler("graphql", "/graphql")
}
