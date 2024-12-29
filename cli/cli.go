package cli

import (
	"flag"
	"fmt"
	"net/url"
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
	// commander address
	Address string

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

	if err := c.validate(); err != nil {
		return err
	}
	return nil
}

func (c *Cli) parse() {
	flag.StringVar(&c.Address, "c", "", "commander address. Example: localhost:3000")
	flag.StringVar(&c.Workdir, "w", ".", "workdir. Example: ./aaa")
	flag.BoolVar(&c.Debug, "debug", false, "debug")
	flag.BoolVar(&c.Version, "version", false, "Print app version")
	flag.Parse()
}

func (c *Cli) IsCommander() bool {
	return c.Address == ""
}

func (c *Cli) validate() error {
	endpoint := fmt.Sprintf("http://%s/graphql", c.Address)
	_, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cli) GetAddress() string {
	return c.Address
}

func (c *Cli) GetWorkdir() string {
	return c.Workdir
}

func (c *Cli) IsDebug() bool {
	return c.Debug
}
