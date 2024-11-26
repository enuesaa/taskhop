package connector

import (
	"context"

	"github.com/enuesaa/taskhop/commander/gql/model"
	"github.com/enuesaa/taskhop/runner/client"
)

func (c *Connector) Run(task client.GetTask_Task) {
	input := model.LogInput{
		Output: "hello",
	}
	c.client.Log(context.Background(), input)
}
