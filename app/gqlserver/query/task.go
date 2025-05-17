package query

import (
	"context"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
	"github.com/enuesaa/taskhop/lib/taskfx"
)

func (r *QueryResolver) Task(ctx context.Context) (*model.Task, error) {
	f := r.Lib.Task.Get()

	var statusgql model.TaskStatus
	switch f.Status {
	case taskfx.StatusRegistration:
		statusgql = model.TaskStatusRegistration
	case taskfx.StatusPrompt:
		statusgql = model.TaskStatusPrompt
	case taskfx.StatusProceeding:
		statusgql = model.TaskStatusProceeding
	}

	res := model.Task{
		Status: statusgql,
		Cmd:    f.Cmd,
	}
	return &res, nil
}
