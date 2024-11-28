package main

import (
	"log"

	"github.com/enuesaa/taskhop/commander"
	"github.com/enuesaa/taskhop/internal/cli"
	"github.com/enuesaa/taskhop/runner"
)

func main() {
	config := cli.Launch()
	if err := config.Validate(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	if config.IsCommander() {
		app := commander.New(config)
		app.Run()
	} else {
		app := runner.New(config)
		app.Run()
	}
}
