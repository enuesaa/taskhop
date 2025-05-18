package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/enuesaa/taskhop/conf"
	flag "github.com/spf13/pflag"
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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "taskhop-agent\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&c.config.VersionFlag, "version", false, "Print version")
	flag.BoolVarP(&c.config.HelpFlag, "help", "h", false, "Print help message")
	flag.Parse()

	if len(os.Args) > 1 {
		c.config.Address = os.Args[1]
	} else {
		c.config.Address = "localhost:3000"
	}
}

func (c *Cli) validate() error {
	endpoint := fmt.Sprintf("http://%s/graphql", c.config.Address)
	if _, err := url.Parse(endpoint); err != nil {
		return err
	}
	return nil
}
