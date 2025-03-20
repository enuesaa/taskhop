package app

import (
	"context"
	"errors"
	"net/http"

	"github.com/enuesaa/taskhop/app/gqlserver"
	"github.com/enuesaa/taskhop/app/gqlserver/middleware"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func NewCommander(config *conf.Config, li lib.Lib, lc fx.Lifecycle, shutdowner fx.Shutdowner) Commander {
	commander := Commander{
		config:     config,
		lib:        li,
		lc:         lc,
		shutdowner: shutdowner,
	}
	return commander
}

type Commander struct {
	config     *conf.Config
	lib        lib.Lib
	lc         fx.Lifecycle
	shutdowner fx.Shutdowner
}

func (a *Commander) Run() error {
	go a.monitor2shutdown()

	server := http.Server{
		Addr:    ":3000",
		Handler: a.handle(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			a.lib.Log.Info(context.Background(), "Error: %s", err.Error())
		}
	}()

	a.lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	return nil
}

func (a *Commander) monitor2shutdown() {
	for range a.lib.Task.SubscribeEnd() {
		a.shutdowner.Shutdown()
	}
}

func (a *Commander) handle() *chi.Mux {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.Recover())
	router.Use(middleware.NoCache())
	router.Use(middleware.Cors())

	// routes
	router.Handle("/graphql", gqlserver.HandleGQL(a.lib))
	router.Handle("/graphql/playground", gqlserver.HandlePlayground())
	router.Handle("/assets", gqlserver.HandleAssets(a.lib))

	return router
}
