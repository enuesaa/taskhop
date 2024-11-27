package main

import (
	"flag"
)

func LaunchCLI() CLI {
	cli := CLI{
		Connect: "",
		Workdir: "",
	}
	flag.StringVar(&cli.Connect, "c", "", "commander address. Example: localhost:3000")
	flag.StringVar(&cli.Workdir, "w", ".", "workdir. Example: ./aaa")
	flag.Parse()

	return cli
}

type CLI struct {
	Connect string
	Workdir string
}

func (c *CLI) IsCommander() bool {
	return c.Connect == ""
}
