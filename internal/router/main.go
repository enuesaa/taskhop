package router

import (
	"github.com/enuesaa/taskhop/internal/repository"
	"github.com/enuesaa/taskhop/internal/routegql"
	"github.com/enuesaa/taskhop/internal/routegqlplayground"
	"github.com/enuesaa/taskhop/ui"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(repos repository.Repos) *echo.Echo {
	app := echo.New()

	// middleware
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// routes
	app.Any("/graphql", routegql.Handle(repos))
	app.GET("/graphql/playground", routegqlplayground.Handle())
	app.Any("/*", ui.Handle())

	return app
}
