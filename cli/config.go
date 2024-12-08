package cli

import (
	"fmt"
	"net/url"
)

type Config struct {
	// commander address
	Address string

	// workdir
	Workdir string
}

func (c *Config) IsCommander() bool {
	return c.Address == ""
}

func (c *Config) Validate() error {
	endpoint := fmt.Sprintf("http://%s/graphql", c.Address)
	_, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	return nil
}
