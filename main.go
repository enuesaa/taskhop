package main

import (
	"net/http"

	"github.com/enuesaa/taskhop/gql"
	"github.com/enuesaa/taskhop/internal/cmd"
	"github.com/enuesaa/taskhop/internal/fs"
	"github.com/enuesaa/taskhop/internal/logging"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			logging.New,
			cmd.New,
			fs.New,
			gql.New,
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
