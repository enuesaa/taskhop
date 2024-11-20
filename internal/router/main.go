package router

import (
	"github.com/enuesaa/taskhop/internal/repository"
	"github.com/enuesaa/taskhop/internal/routegql"
	"github.com/enuesaa/taskhop/internal/routegqlplayground"

	"github.com/enuesaa/taskhop/ui"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func New(repos repository.Repos) *chi.Mux {
	app := chi.NewRouter()

	// middleware
	app.Use(middleware.Logger)
	app.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	// routes
	app.HandleFunc("/graphql", routegql.Handle(repos))
	app.Get("/graphql/playground", routegqlplayground.Handle())
	app.HandleFunc("/*", ui.Handle())

	return app
}
