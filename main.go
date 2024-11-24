package main

import (
	"net/http"

	"github.com/enuesaa/taskhop/gql"
	"github.com/enuesaa/taskhop/internal/cmdexec"
	"github.com/enuesaa/taskhop/internal/fs"
	"github.com/enuesaa/taskhop/internal/logging"
	"github.com/enuesaa/taskhop/internal/usecase"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			logging.New,
			cmdexec.New,
			fs.New,
			gql.New,
			usecase.New,
			NewHandler,
			NewServer,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()

	// ctx := context.Background()
	// ctx = repos.Log.Use(ctx, "a", "b")
	// repos.Log.Info(ctx, "aa")

	// if err := http.ListenAndServe(":3000", router); err != nil {
	// 	log.Fatalf("Error: %s", err.Error())
	// }
}
