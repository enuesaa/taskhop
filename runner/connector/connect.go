package connector

func (c *Connector) Connect(address string) error {
	if err := c.Polling(address); err != nil {
		return err
	}
	if err := c.CheckHealth(address); err != nil {
		return err
	}
	return nil
}
