package connector

import (
	"context"
	"fmt"
)

func (c *Connector) checkHealth() error {
	data, err := c.client.GetHealth(context.Background())
	if err != nil {
		return err
	}

	if !data.Health.Ok {
		return fmt.Errorf("commander not healthy")
	}
	return nil
}
