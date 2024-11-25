package main

import (
	"github.com/enuesaa/taskhop/app"
)

func main() {
	cli := LaunchCLI()

	if cli.IsCommander() {
		commander := app.NewCommander()
		commander.Run()
	} else {
		runner := app.NewRunner()
		runner.Run()
	}
}
