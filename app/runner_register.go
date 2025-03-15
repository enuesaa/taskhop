package app

import (
	"context"
	"errors"

	"github.com/enuesaa/taskhop/app/gql/model"
)

var ErrTaskNotAvailable = errors.New("task not available")

func (a *Runner) Register() error {
	taskres, err := a.conn.GetTask(context.Background())
	if err != nil {
		return err
	}
	task := taskres.Task

	if task.Status != model.TaskStatusRegistration {
		return ErrTaskNotAvailable
	}

	regires, err := a.conn.Register(context.Background())
	if err != nil {
		return err
	}
	if !regires.Register {
		return err
	}
	return nil
}
