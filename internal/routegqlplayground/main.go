package routegqlplayground

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

func Handle() echo.HandlerFunc {
	return echo.WrapHandler(playground.Handler("graphql", "/graphql"))
}
