package usecase

import (
	"context"
	"errors"

	"github.com/enuesaa/taskhop/commander/gql/model"
	"github.com/enuesaa/taskhop/runner/connector"
)

var ErrTaskNotAvailable = errors.New("task not available")

func (u *UseCase) Register() (connector.GetTask_Task, error) {
	u.Logi.Info(context.Background(), "receive a task")
	taskres, err := u.conn.GetTask(context.Background())
	if err != nil {
		return connector.GetTask_Task{}, err
	}
	task := taskres.Task

	if task.Status != model.TaskStatusWaiting {
		return connector.GetTask_Task{}, ErrTaskNotAvailable
	}

	u.Logi.Info(context.Background(), "assign me")
	regires, err := u.conn.Register(context.Background())
	if err != nil {
		return connector.GetTask_Task{}, err
	}
	if !regires.Register {
		return connector.GetTask_Task{}, err
	}

	return task, nil
}
