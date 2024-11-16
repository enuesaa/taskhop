package main

import (
	"log"

	"github.com/enuesaa/taskrun/internal/repository"
	"github.com/enuesaa/taskrun/internal/router"
)

func main() {
	repos := repository.New()
	app := router.New(repos)

	if err := app.Start(":3000"); err != nil {
		log.Fatal(err)
	}
}
