package main

import (
	"context"
	"log"
	"net/http"

	"github.com/enuesaa/taskhop/internal/repository"
)

func main() {
	repos := repository.New()
	router := NewRouter(repos)

	ctx := context.Background()

	ctx = repos.Log.Use(ctx, "a", "b")
	repos.Log.Info(ctx, "aa")

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
