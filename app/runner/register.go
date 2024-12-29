package runner

import (
	"context"
	"errors"

	"github.com/enuesaa/taskhop/app/commander/gql/model"
	"github.com/enuesaa/taskhop/app/runner/connector"
)

var ErrTaskNotAvailable = errors.New("task not available")

func (a *App) Register() (connector.GetTask_Task, error) {
	taskres, err := a.conn.GetTask(context.Background())
	if err != nil {
		return connector.GetTask_Task{}, err
	}
	task := taskres.Task

	if task.Status != model.TaskStatusWaiting {
		return connector.GetTask_Task{}, ErrTaskNotAvailable
	}

	regires, err := a.conn.Register(context.Background())
	if err != nil {
		return connector.GetTask_Task{}, err
	}
	if !regires.Register {
		return connector.GetTask_Task{}, err
	}

	return task, nil
}
