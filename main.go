package main

import (
	"log"
	"net/http"

	"github.com/enuesaa/taskhop/internal/repository"
	"github.com/enuesaa/taskhop/internal/router"
)

func main() {
	repos := repository.New()
	app := router.New(repos)

	if err := http.ListenAndServe(":3000", app); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
