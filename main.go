package main

import (
	// "log"
	// "net/http"

	// "github.com/enuesaa/taskhop/internal/repository"
	"net/http"

	"github.com/enuesaa/taskhop/internal/repository"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			NewServer,
			NewHandler,
			repository.New,
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
