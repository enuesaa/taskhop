package main

import (
	"net/http"

	"github.com/enuesaa/taskhop/gql"

	"github.com/enuesaa/taskhop/ui"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewServer(gqhandle http.HandlerFunc) *http.Server {
	router := chi.NewRouter()

	// middleware
	// app.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		start := time.Now()
	// 		next.ServeHTTP(w, r)
	// 		log.Info(r.Context(), "%s %s %s", r.Method, r.URL.Path, time.Since(start))
	// 	})
	// })
	router.Use(middleware.Recoverer)
	router.Use(middleware.NoCache)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	// routes
	router.HandleFunc("/graphql", gqhandle)
	router.Get("/graphql/playground", gql.HandlePlayground())
	router.HandleFunc("/*", ui.Handle())

	srv := &http.Server{
		Addr: ":3000",
		Handler: router,
	}
	return srv
}
