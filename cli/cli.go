package cli

import (
	"flag"
	"fmt"
	"net/url"
)

type Cli struct {
	// commander address
	Address string

	// workdir
	Workdir string
}

func (c *Cli) Launch() error {
	c.parse()

	if err := c.validate(); err != nil {
		return err
	}
	return nil
}

func (c *Cli) parse() {
	flag.StringVar(&c.Address, "c", "", "commander address. Example: localhost:3000")
	flag.StringVar(&c.Workdir, "w", ".", "workdir. Example: ./aaa")
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
