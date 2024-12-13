package commander

import (
	"context"
	"errors"
	"net/http"

	"github.com/enuesaa/taskhop/app/commander/gql"
	"github.com/enuesaa/taskhop/app/commander/gqlplayground"
	"github.com/enuesaa/taskhop/app/commander/storage"
	"github.com/enuesaa/taskhop/app/commander/middleware"
	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func New(cl cli.ICli, li lib.Lib, lc fx.Lifecycle) App {
	app := App{
		cli: cl,
		lib: li,
		lc: lc,
	}
	return app
}

type App struct {
	cli cli.ICli
	lib lib.Lib
	lc fx.Lifecycle
}

func (a *App) Run() error {
	if err := a.load(); err != nil {
		return err
	}
	if err := a.start(); err != nil {
		return err
	}
	return nil
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
	a.lib.Log.Info(context.Background(), "started")
	return nil
}

func (a *App) start() error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router(),
	}
	a.lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil {
					if errors.Is(err, http.ErrServerClosed) {
						return
					}
					a.lib.Log.Info(context.Background(), "Error: %s", err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			a.lib.Log.Info(context.Background(), "stop app")
			return server.Shutdown(ctx)
		},
	})
	return nil
}

func (a *App) router() *chi.Mux {
	router := chi.NewRouter()

	// middleware
	// router.Use(middleware.Logger(c))
	router.Use(middleware.Recover())
	router.Use(middleware.NoCache())
	router.Use(middleware.Cors())

	// routes
	router.Handle("/graphql", gql.Handle(a.lib))
	router.Handle("/graphql/playground", gqlplayground.Handle())
	router.Handle("/storage/archive", storage.Handle(a.lib))

	return router
}