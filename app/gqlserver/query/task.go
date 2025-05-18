package query

import (
	"context"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
	"github.com/enuesaa/taskhop/lib/taskfx"
)

func (r *QueryResolver) Task(ctx context.Context) (*model.Task, error) {
	f := r.li.Task.Get()

	var status model.TaskStatus
	switch f.Status {
	case taskfx.StatusRegistration:
		status = model.TaskStatusRegistration
	case taskfx.StatusPrompt:
		status = model.TaskStatusPrompt
	case taskfx.StatusProceeding:
		status = model.TaskStatusProceeding
	}

	if f.Text == "@push" {
		res := model.Task{
			Status: status,
			Method: model.TaskMethodDownloadAsset,
		}
		return &res, nil
	}
	if f.Text == "@pull" {
		res := model.Task{
			Status: status,
			Method: model.TaskMethodUploadAsset,
		}
		return &res, nil
	}

	res := model.Task{
		Status: status,
		Method: model.TaskMethodCmd,
		Cmd:    f.Text,
	}
	return &res, nil
}
