package commander

import (
	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib"
	"go.uber.org/fx"
)

func New(cl cli.ICli, li lib.Lib, lc fx.Lifecycle) IApp {
	app := App{
		cli: cl,
		lib: li,
		lc: lc,
	}
	return &app
}

// func NewOld(config cli.Config) *fx.App {
// 	// see https://github.com/ankorstore/yokai
// 	app := fx.New(
// 		internal.Module,
// 		fx.Supply(config),
// 		fx.Provide(NewRouter),
// 		fx.Invoke(func(c internal.Container) error {
// 			task, err := c.Taski.Read()
// 			if err != nil {
// 				c.Logi.Info(context.Background(), "Error: %s", err.Error())
// 				return err
// 			}
// 			if err := c.Taski.Validate(task); err != nil {
// 				c.Logi.Info(context.Background(), "Error: %s", err.Error())
// 				return err
// 			}
// 			c.Logi.Info(context.Background(), "started")
// 			return nil
// 		}),
// 		// fx.Invoke(func(c internal.Container, shutdowner fx.Shutdowner) error {
// 		// 	go func() {
// 		// 		for status := range c.Taski.Subscribe() {
// 		// 			if status == taskfx.StatusCompleted {
// 		// 				shutdowner.Shutdown()
// 		// 				break
// 		// 			}
// 		// 		}
// 		// 	}()
// 		// 	return nil
// 		// }),
// 		fx.Invoke(func(lc fx.Lifecycle, router *chi.Mux, c internal.Container) {
// 			server := &http.Server{
// 				Addr:    ":3000",
// 				Handler: router,
// 			}
// 			lc.Append(fx.Hook{
// 				OnStart: func(ctx context.Context) error {
// 					go func() {
// 						if err := server.ListenAndServe(); err != nil {
// 							if errors.Is(err, http.ErrServerClosed) {
// 								return
// 							}
// 							c.Logi.Info(context.Background(), "Error: %s", err.Error())
// 						}
// 					}()
// 					return nil
// 				},
// 				OnStop: func(ctx context.Context) error {
// 					c.Logi.Info(context.Background(), "stop app")
// 					return server.Shutdown(ctx)
// 				},
// 			})
// 		}),
// 		// fx.NopLogger,
// 	)

// 	return app
// }
