package main

import (
	"fmt"
	"net/url"
	"os"

	flag "github.com/spf13/pflag"
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
	if c.config.HelpFlag {
		flag.Usage()
		os.Exit(0)
	}
	if err := c.validate(); err != nil {
		return err
	}
	return nil
}

func (c *Cli) parse() {
	flag.StringVarP(&c.config.Address, "connect", "c", "localhost:3000", "Commander address")
	flag.BoolVar(&c.config.VersionFlag, "version", false, "Print taskhop version")
	flag.BoolVarP(&c.config.HelpFlag, "help", "h", false, "Print help message")
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
