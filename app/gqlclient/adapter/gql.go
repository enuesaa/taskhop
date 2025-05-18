package adapter

import (
	"context"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
)

func (c *Adapter) Log(output string) error {
	input := model.LogInput{
		Output: output,
	}
	_, err := c.gql.Log(context.Background(), input)
	if err != nil {
		return err
	}
	return nil
}

func (c *Adapter) Completed() error {
	if _, err := c.gql.Completed(context.Background()); err != nil {
		return err
	}
	return nil
}

func (c *Adapter) GetTask() (*GetTask, error) {
	return c.gql.GetTask(context.Background())
}

func (c *Adapter) GetHealth() error {
	if _, err := c.gql.GetHealth(context.Background()); err != nil {
		return err
	}
	return nil
}

func (c *Adapter) Register() error {
	if _, err := c.gql.Register(context.Background()); err != nil {
		return err
	}
	return nil
}
