package connector

import (
	"context"

	"github.com/enuesaa/taskhop/commander/gql/model"
	"github.com/enuesaa/taskhop/runner/client"
)

func (c *Connector) Run(task client.GetTask_Task) error {
	w := LogWriter{
		Callback: func (data string) error {
			input := model.LogInput{
				Output: data,
			}
			_, err := c.client.Log(context.Background(), input)
			if err != nil {
				return err
			}
			return nil
		},
	}

	if err := c.Cmdi.MultiExec(&w, task.Cmds); err != nil {
		return err
	}

	return nil
}
