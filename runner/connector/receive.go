package connector

import (
	"context"
	"errors"

	"github.com/enuesaa/taskhop/commander/gql/model"
	"github.com/enuesaa/taskhop/runner/client"
)

var ErrTaskNotAvailable = errors.New("tasl not available")

func (c *Connector) Receive() (client.GetTask_Task, error) {
	data, err := c.client.GetTask(context.Background())
	if err != nil {
		return client.GetTask_Task{}, err
	}
	task := data.Task

	if task.Status != model.TaskStatusWaiting {
		return client.GetTask_Task{}, ErrTaskNotAvailable
	}
	return task, nil
}
