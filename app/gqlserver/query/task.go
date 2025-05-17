package query

import (
	"context"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
	"github.com/enuesaa/taskhop/lib/taskfx"
)

func (r *QueryResolver) Task(ctx context.Context) (*model.Task, error) {
	f := r.Lib.Task.Get()

	var status model.TaskStatus
	switch f.Status {
	case taskfx.StatusRegistration:
		status = model.TaskStatusRegistration
	case taskfx.StatusPrompt:
		status = model.TaskStatusPrompt
	case taskfx.StatusProceeding:
		status = model.TaskStatusProceeding
	}

	method := model.TaskMethodCmd
	cmd := f.Text
	if cmd == "@push" {
		method = model.TaskMethodDownloadAsset
		cmd = ""
	}

	res := model.Task{
		Status: status,
		Method: method,
		Cmd:    cmd,
	}
	return &res, nil
}
