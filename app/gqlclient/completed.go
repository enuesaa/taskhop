package gqlclient

import "context"

func (c *Connector) MarkCompleted(ctx context.Context) error {
	if _, err := c.gql.Completed(ctx); err != nil {
		return err
	}
	return nil
}
