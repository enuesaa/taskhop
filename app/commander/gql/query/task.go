package query

import (
	"context"

	"github.com/enuesaa/taskhop/app/commander/gql/model"
	"github.com/enuesaa/taskhop/lib/taskfx"
)

func (r *QueryResolver) Task(ctx context.Context) (*model.Task, error) {
	f, err := r.Lib.Task.Read()
	if err != nil {
		return nil, err
	}

	status := r.Lib.Task.GetStatus()

	var statusgql model.TaskStatus
	switch status {
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
