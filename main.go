package main

import (
	"log"

	"github.com/enuesaa/taskhop/internal/repository"
	"github.com/enuesaa/taskhop/internal/router"
)

func main() {
	repos := repository.New()
	app := router.New(repos)

	if err := app.Start(":3000"); err != nil {
		log.Fatal(err)
	}
}
