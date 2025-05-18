package adapter

import (
	"context"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
)

func (c *Adapter) Log(ctx context.Context, output string) error {
	input := model.LogInput{
		Output: output,
	}
	_, err := c.gql.Log(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (c *Adapter) Completed(ctx context.Context) error {
	if _, err := c.gql.Completed(ctx); err != nil {
		return err
	}
	return nil
}

func (c *Adapter) GetTask(ctx context.Context) (*GetTask, error) {
	return c.gql.GetTask(ctx)
}

func (c *Adapter) GetHealth(ctx context.Context) error {
	if _, err := c.gql.GetHealth(ctx); err != nil {
		return err
	}
	return nil
}

func (c *Adapter) Register(ctx context.Context) error {
	if _, err := c.gql.Register(ctx); err != nil {
		return err
	}
	return nil
}
