package connector

func (c *Connector) Connect(commanderAddress string) error {
	if err := c.polling(commanderAddress); err != nil {
		return err
	}
	if err := c.checkHealth(); err != nil {
		return err
	}
	return nil
}
