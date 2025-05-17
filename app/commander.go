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

func NewCommander(config *conf.Config, li lib.Lib, lc fx.Lifecycle, shutdowner fx.Shutdowner) *Commander {
	commander := &Commander{
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
	server     *http.Server
}

func (a *Commander) Run(ctx context.Context) error {
	a.server = &http.Server{
		Addr:    ":3000",
		Handler: a.handle(),
	}
	go a.listen(ctx)
	go a.monitor(ctx)

	a.lc.Append(fx.Hook{
		OnStop: a.close,
	})
	return nil
}

func (a *Commander) listen(ctx context.Context) {
	if err := a.server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return
		}
		a.lib.Log.Info(ctx, "Error: %s", err.Error())
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
	router.Handle("/upload", gqlserver.HandleUpload(a.lib, a.config))

	return router
}

func (a *Commander) monitor(ctx context.Context) {
	for err := range a.lib.Task.Subscribe() {
		a.lib.Log.Info(ctx, "Error: %s", err.Error())
		a.shutdowner.Shutdown()
	}
}

func (a *Commander) close(ctx context.Context) error {
	if a.server != nil {
		a.server.Shutdown(ctx)
	}
	return nil
}
