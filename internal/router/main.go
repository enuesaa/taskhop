package router

import (
	"github.com/enuesaa/taskhop/internal/repository"
	"github.com/enuesaa/taskhop/internal/routegql"
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

	app.HideBanner = true
	app.HidePort = true

	return app
}
