package connector

import (
	"context"
	"errors"

	"github.com/enuesaa/taskhop/commander/gql/model"
	"github.com/enuesaa/taskhop/runner/client"
)

var ErrTaskNotAvailable = errors.New("task not available")

func (c *Connector) Register() (client.GetTask_Task, error) {
	taskres, err := c.client.GetTask(context.Background())
	if err != nil {
		return client.GetTask_Task{}, err
	}
	task := taskres.Task

	if task.Status != model.TaskStatusWaiting {
		return client.GetTask_Task{}, ErrTaskNotAvailable
	}

	regires, err := c.client.Register(context.Background())
	if err != nil {
		return client.GetTask_Task{}, err
	}
	if !regires.Register {
		return client.GetTask_Task{}, err
	}

	return task, nil
}
