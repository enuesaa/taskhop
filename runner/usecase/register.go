package usecase

import (
	"context"
	"errors"

	"github.com/enuesaa/taskhop/commander/gql/model"
	"github.com/enuesaa/taskhop/runner/gqlclient"
)

var ErrTaskNotAvailable = errors.New("task not available")

func (c *UseCase) Register() (gqlclient.GetTask_Task, error) {
	taskres, err := c.client.GetTask(context.Background())
	if err != nil {
		return gqlclient.GetTask_Task{}, err
	}
	task := taskres.Task

	if task.Status != model.TaskStatusWaiting {
		return gqlclient.GetTask_Task{}, ErrTaskNotAvailable
	}

	regires, err := c.client.Register(context.Background())
	if err != nil {
		return gqlclient.GetTask_Task{}, err
	}
	if !regires.Register {
		return gqlclient.GetTask_Task{}, err
	}

	return task, nil
}
