package runner

func New() IApp {
	return &App{}
}

// func New(config cli.Config) *fx.App {
// 	app := fx.New(
// 		internal.Module,
// 		fx.Supply(config),
// 		fx.Provide(connector.New, usecase.New),
// 		fx.Invoke(func (u usecase.UseCase) error {
// 			for {
// 				if err := u.Connect(); err != nil {
// 					return err
// 				}
// 				task, err := u.Register()
// 				if err != nil {
// 					return err
// 				}
			
// 				if err := u.UnArchive(); err != nil {
// 					return err
// 				}
// 				if err := u.Run(task); err != nil {
// 					return err
// 				}
// 			}
// 		}),
// 		fx.NopLogger,
// 	)

// 	return app
// }
