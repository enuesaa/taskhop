package commander

import (
	"github.com/enuesaa/taskhop/commander/gql"
	"github.com/enuesaa/taskhop/commander/gqlplayground"
	"github.com/enuesaa/taskhop/commander/middleware"
	"github.com/enuesaa/taskhop/commander/storage"
	"github.com/enuesaa/taskhop/internal"
	"github.com/go-chi/chi/v5"
)

func NewRouter(c internal.Container) *chi.Mux {
	router := chi.NewRouter()

	// middleware
	// router.Use(middleware.Logger(c))
	router.Use(middleware.Recover())
	router.Use(middleware.NoCache())
	router.Use(middleware.Cors())

	// routes
	router.Handle("/graphql", gql.Handle(c))
	router.Handle("/graphql/playground", gqlplayground.Handle())
	router.Handle("/storage/archive", storage.Handle(c))

	return router
}
