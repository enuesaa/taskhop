package router

import (
	"log/slog"
	"time"

	"github.com/enuesaa/taskhop/internal/repository"
	"github.com/enuesaa/taskhop/internal/routegql"
	"github.com/enuesaa/taskhop/internal/routegqlplayground"

	"github.com/enuesaa/taskhop/ui"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog/v2"
)

func New(repos repository.Repos) *chi.Mux {
	app := chi.NewRouter()

	// middleware
	logger := httplog.NewLogger("http", httplog.Options{
		LogLevel:        slog.LevelDebug,
		TimeFieldFormat: time.RFC3339,
		// Concise:          true,
		// RequestHeaders:   true,
	})
	app.Use(httplog.RequestLogger(logger))
	app.Use(middleware.Recoverer)
	app.Use(middleware.NoCache)
	app.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	// routes
	app.HandleFunc("/graphql", routegql.Handle(repos))
	app.Get("/graphql/playground", routegqlplayground.Handle())
	app.HandleFunc("/*", ui.Handle())

	return app
}
