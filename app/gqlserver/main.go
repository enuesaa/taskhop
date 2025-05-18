package gqlserver

import (
	"github.com/enuesaa/taskhop/app/gqlserver/middleware"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
	"github.com/go-chi/chi/v5"
)

func New(config *conf.Config, li *lib.Lib) *chi.Mux {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.Recover())
	router.Use(middleware.NoCache())
	router.Use(middleware.Cors())

	// routes
	handler := Handler{
		li:     li,
		config: config,
	}
	router.Handle("/graphql", handler.HandleGQL())
	router.Handle("/graphql/playground", handler.HandlePlayground())
	router.Get("/assets", handler.Assets)
	router.Post("/upload", handler.Upload)

	return router
}

type Handler struct {
	li     *lib.Lib
	config *conf.Config
}
