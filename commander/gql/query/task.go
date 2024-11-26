package query

import (
	"context"

	"github.com/enuesaa/taskhop/commander/gql/model"
	"github.com/enuesaa/taskhop/internal/runfx"
)

func (r *QueryResolver) Task(ctx context.Context) (*model.Task, error) {
	f, err := r.Taski.Read("testdata/cmds.yml")
	if err != nil {
		return nil, err
	}

	status := r.Runi.GetStatus()

	var statusgql model.TaskStatus
	switch status {
	case runfx.StatusWaiting:
		statusgql = model.TaskStatusWaiting
	case runfx.StatusProceeding:
		statusgql = model.TaskStatusProceeding
	case runfx.StatusCompleted:
		statusgql = model.TaskStatusCompleted
	}

	res := model.Task{
		Status: statusgql,
		Cmds: f.Cmds,
	}
	return &res, nil
}
