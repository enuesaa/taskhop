package cli

type Config struct {
	// commander address
	Address string

	// workdir
	Workdir string
}

func (c *Config) IsCommander() bool {
	return c.Address == ""
}
