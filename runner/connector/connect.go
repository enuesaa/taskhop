package connector

func (c *Connector) Connect() error {
	if err := c.polling(); err != nil {
		return err
	}
	if err := c.checkHealth(); err != nil {
		return err
	}
	return nil
}
