package main

import (
	"github.com/enuesaa/taskhop/commander"
	"github.com/enuesaa/taskhop/runner"
)

func main() {
	cli := LaunchCLI()

	if cli.IsCommander() {
		app := commander.New(cli.Workdir)
		app.Run()
	} else {
		app := runner.New(cli.Workdir, cli.Connect)
		app.Run()
	}
}
