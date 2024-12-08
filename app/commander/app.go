package commander

import (
	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib"
	"go.uber.org/fx"
)

type IApp interface {
	Run() error
}

type App struct {
	cli cli.ICli
	lib lib.ILib
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
	task, err := a.Taski.Read()
	if err != nil {
		c.Logi.Info(context.Background(), "Error: %s", err.Error())
		return err
	}
	if err := c.Taski.Validate(task); err != nil {
		c.Logi.Info(context.Background(), "Error: %s", err.Error())
		return err
	}
	c.Logi.Info(context.Background(), "started")
	return nil
}

func (a *App) start() error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil {
					if errors.Is(err, http.ErrServerClosed) {
						return
					}
					c.Logi.Info(context.Background(), "Error: %s", err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			c.Logi.Info(context.Background(), "stop app")
			return server.Shutdown(ctx)
		},
	})
	return nil
}
