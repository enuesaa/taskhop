package main

import (
	"flag"
)

func LaunchCLI() CLI {
	cli := CLI{
		Commander: "",
	}
	flag.StringVar(&cli.Commander, "commander", "", "commander address. Example: localhost:3000")
	flag.Parse()

	return cli
}

type CLI struct {
	Commander string
}

func (c *CLI) IsCommander() bool {
	return c.Commander == ""
}
