package gqlclient

func (c *Connector) Log(text string) {
	c.LogWriter.Write([]byte(text))
}
