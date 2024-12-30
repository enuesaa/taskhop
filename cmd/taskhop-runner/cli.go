package main

import (
	"flag"
	"fmt"
	"net/url"
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
	if err := c.validate(); err != nil {
		return err
	}
	return nil
}

func (c *Cli) parse() {
	flag.StringVar(&c.config.Address, "c", "localhost:3000", "Commander address")
	flag.BoolVar(&c.config.VersionFlag, "version", false, "Print taskhop version")
	flag.Parse()
}

func (c *Cli) validate() error {
	endpoint := fmt.Sprintf("http://%s/graphql", c.config.Address)
	_, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	return nil
}
