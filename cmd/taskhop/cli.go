package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/enuesaa/taskhop/conf"
)

func New(config *conf.Config) ICli {
	cli := Cli{
		config: config,
	}
	return &cli
}

type ICli interface {
	Launch() error
}

type Cli struct {
	config *conf.Config
}

func (c *Cli) Launch() error {
	c.parse()

	if c.config.VersionFlag {
		fmt.Printf("%s\n", c.config.Version)
		os.Exit(0)
	}
	return nil
}

func (c *Cli) parse() {
	flag.StringVar(&c.config.Workdir, "w", ".", "workdir. Example: ./aaa")
	flag.BoolVar(&c.config.VersionFlag, "version", false, "Print taskhop version")
	flag.Parse()
}
