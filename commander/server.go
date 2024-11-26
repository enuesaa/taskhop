package commander

import (
	"net/http"
	"time"

	"github.com/enuesaa/taskhop/commander/gql"
	"github.com/enuesaa/taskhop/commander/gqlplayground"
	"github.com/enuesaa/taskhop/internal/logging"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewServer(logi logging.I) *http.Server {
	router := chi.NewRouter()

	// middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logi.Info(r.Context(), "%s %s %s", r.Method, r.URL.Path, time.Since(start))
		})
	})
	router.Use(middleware.Recoverer)
	router.Use(middleware.NoCache)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	// routes
	router.HandleFunc("/graphql", gql.Handle())
	router.Get("/graphql/playground", gqlplayground.Handle())

	srv := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	return srv
}
