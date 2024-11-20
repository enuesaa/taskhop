package main

import (
	"context"
	"log"
	"net/http"

	"github.com/enuesaa/taskhop/internal/repository"
	"github.com/enuesaa/taskhop/internal/router"
)

func main() {
	repos := repository.New()
	app := router.New(repos)

	ctx := context.Background()

	ctx = repos.Log.Use(ctx, "a", "b")
	repos.Log.Info(ctx, "aa")

	if err := http.ListenAndServe(":3000", app); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
