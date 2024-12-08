package query

import (
	"context"

	"github.com/enuesaa/taskhop/commander/gql/model"
	"github.com/enuesaa/taskhop/internal/taskfx"
)

func (r *QueryResolver) Task(ctx context.Context) (*model.Task, error) {
	f, err := r.Taski.Read()
	if err != nil {
		return nil, err
	}

	status := r.Taski.GetStatus()

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
		Cmds: f.Cmds,
	}
	return &res, nil
}
