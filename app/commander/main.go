package commander

import (
	"context"
	"errors"
	"net/http"

	"github.com/enuesaa/taskhop/app/commander/gql"
	"github.com/enuesaa/taskhop/app/commander/gqlplayground"
	"github.com/enuesaa/taskhop/app/commander/middleware"
	"github.com/enuesaa/taskhop/app/commander/storage"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func New(config conf.Config, li lib.Lib, lc fx.Lifecycle, shutdowner fx.Shutdowner) App {
	app := App{
		config:     config,
		lib:        li,
		lc:         lc,
		shutdowner: shutdowner,
	}
	return app
}

type App struct {
	config     conf.Config
	lib        lib.Lib
	lc         fx.Lifecycle
	shutdowner fx.Shutdowner
}

func (a *App) Run() error {
	go a.monitor2shutdown()

	if err := a.load(); err != nil {
		return err
	}
	if err := a.serve(); err != nil {
		return err
	}

	return nil
}

func (a *App) monitor2shutdown() {
	for range a.lib.Proc.SubscribeCompleted() {
		a.shutdowner.Shutdown()
	}
}

func (a *App) load() error {
	task, err := a.lib.Task.Read()
	if err != nil {
		a.lib.Log.Info(context.Background(), "Error: %s", err.Error())
		return err
	}
	if err := a.lib.Task.Validate(task); err != nil {
		a.lib.Log.Info(context.Background(), "Error: %s", err.Error())
		return err
	}
	return nil
}

func (a *App) serve() error {
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

func (a *App) handle() *chi.Mux {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.Recover())
	router.Use(middleware.NoCache())
	router.Use(middleware.Cors())

	// routes
	router.Handle("/graphql", gql.Handle(a.lib))
	router.Handle("/graphql/playground", gqlplayground.Handle())
	router.Handle("/storage/archive", storage.Handle(a.lib))

	return router
}
