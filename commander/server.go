package commander

import (
	"net/http"

	"github.com/enuesaa/taskhop/commander/middleware"
	"github.com/enuesaa/taskhop/internal"
	"github.com/go-chi/chi/v5"
)

func NewServer(routes []Route, c internal.Container) *http.Server {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.Logger(c))
	router.Use(middleware.Recover())
	router.Use(middleware.NoCache())
	router.Use(middleware.Cors())

	// routes
	for _, route := range routes {
		router.Handle(route.Pattern(), route)
	}

	srv := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	return srv
}
