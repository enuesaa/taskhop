package app

import (
	"context"
	"errors"

	"github.com/enuesaa/taskhop/app/gql/model"
)

var ErrTaskNotAvailable = errors.New("task not available")

func (a *Runner) Connect() error {
	a.lib.Log.AppInfo(context.Background(), "polling...")

	if err := a.conn.DialPolling(); err != nil {
		return err
	}
	if _, err := a.conn.GetHealth(context.Background()); err != nil {
		return err
	}
	return nil
}

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
		return ErrTaskNotAvailable
	}
	return nil
}
