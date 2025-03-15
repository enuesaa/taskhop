package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/enuesaa/taskhop/app/gql"
	"github.com/enuesaa/taskhop/app/gqlplayground"
	"github.com/enuesaa/taskhop/app/middleware"
	"github.com/enuesaa/taskhop/app/storage"
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

	text, err := Prompt()
	if err != nil {
		return err
	}
	fmt.Printf("input: %s\n", text)

	if err := a.load(); err != nil {
		return err
	}
	if err := a.serve(); err != nil {
		return err
	}
	return nil
}

func (a *Commander) monitor2shutdown() {
	for range a.lib.Proc.SubscribeCompleted() {
		a.shutdowner.Shutdown()
	}
}

func (a *Commander) load() error {
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

func (a *Commander) serve() error {
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

func (a *Commander) handle() *chi.Mux {
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
