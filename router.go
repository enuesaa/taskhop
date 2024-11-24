package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/enuesaa/taskhop/gql"
	"github.com/enuesaa/taskhop/internal/repository"
	"go.uber.org/fx"

	"github.com/enuesaa/taskhop/ui"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewHandler(repos repository.Repos) http.Handler {
	app := chi.NewRouter()

	// middleware
	app.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			repos.Log.Info(r.Context(), "%s %s %s", r.Method, r.URL.Path, time.Since(start))
		})
	})
	app.Use(middleware.Recoverer)
	app.Use(middleware.NoCache)
	app.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	// routes
	app.HandleFunc("/graphql", gql.Handle(repos))
	app.Get("/graphql/playground", gql.HandlePlayground())
	app.HandleFunc("/*", ui.Handle())

	return app
}

func NewServer(lc fx.Lifecycle, handler http.Handler) *http.Server {
	srv := &http.Server{
		Addr: ":3000",
		Handler: handler,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting HTTP server")
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
