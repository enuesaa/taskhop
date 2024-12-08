package cli


type CliMock struct {
	Cli
}

func (c *CliMock) Launch() error {
	// As this struct is mock, do not run c.parse() here.

	if err := c.validate(); err != nil {
		return err
	}
	return nil
}
