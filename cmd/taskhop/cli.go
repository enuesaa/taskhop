package main

import (
	"flag"
	"fmt"
	"os"
)

type ICli interface {
	Launch() error
	IsCommander() bool
	GetAddress() string
	GetWorkdir() string
	IsDebug() bool
}

type Cli struct {
	// workdir
	Workdir string

	// debug
	Debug bool

	// version
	Version bool
}

func (c *Cli) Launch() error {
	c.parse()

	if c.Version {
		fmt.Printf("0.0.3\n")
		os.Exit(0)
	}
	return nil
}

func (c *Cli) parse() {
	flag.StringVar(&c.Workdir, "w", ".", "workdir. Example: ./aaa")
	flag.BoolVar(&c.Version, "version", false, "Print app version")
	flag.Parse()
}

func (c *Cli) GetWorkdir() string {
	return c.Workdir
}

func (c *Cli) IsDebug() bool {
	return c.Debug
}
