package query

import (
	"context"

	"github.com/enuesaa/taskhop/app/gql/model"
	"github.com/enuesaa/taskhop/lib/taskfx"
)

func (r *QueryResolver) Task(ctx context.Context) (*model.Task, error) {
	f, err := r.Lib.Task.Get()
	if err != nil {
		return nil, err
	}

	var statusgql model.TaskStatus
	switch f.Status {
	case taskfx.StatusWaiting:
		statusgql = model.TaskStatusWaiting
	case taskfx.StatusProceeding:
		statusgql = model.TaskStatusProceeding
	case taskfx.StatusCompleted:
		statusgql = model.TaskStatusCompleted
	}

	res := model.Task{
		Status: statusgql,
		Cmds:   f.Cmds,
	}
	return &res, nil
}
