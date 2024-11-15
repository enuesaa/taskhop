package router

import (
	"github.com/enuesaa/taskrun/internal/repository"
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

	app.HideBanner = true
	app.HidePort = true

	return app
}
