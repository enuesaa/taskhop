package gqlclient

import "context"

func (c *Connector) MarkCompleted(ctx context.Context) error {
	_, err := c.gql.Completed(ctx)
	if err != nil {
		return err
	}
	return nil
}
