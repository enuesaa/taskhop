package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/enuesaa/taskhop/conf"
)

type ICli interface {
	Launch() error
	GetConfig() conf.Config
}

type Cli struct {
	// commander address
	Address string

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
	flag.BoolVar(&c.Version, "version", false, "Print app version")
	flag.Parse()
}

func (c *Cli) validate() error {
	endpoint := fmt.Sprintf("http://%s/graphql", c.Address)
	_, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	return nil
}
