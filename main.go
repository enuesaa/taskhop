package main

import (
	"github.com/enuesaa/taskhop/commander"
	"github.com/enuesaa/taskhop/internal/cli"
	"github.com/enuesaa/taskhop/runner"
)

func main() {
	config := cli.Launch()

	if config.IsCommander() {
		app := commander.New(config)
		app.Run()
	} else {
		app := runner.New(config)
		app.Run()
	}
}
