package gqlclient

import (
	"context"
	"errors"
	"iter"
	"time"

	"github.com/enuesaa/taskhop/app/gqlclient/adapter"
	"github.com/enuesaa/taskhop/app/gqlserver/model"
)

var ErrConnectOnGetTask = errors.New("failed to get task")

type Task struct {
	Err error
	res *adapter.GetTask
}

func (t *Task) IsDownload() bool {
	return t.res.Task.Method == model.TaskMethodDownloadAsset
}

func (t *Task) IsUpload() bool {
	return t.res.Task.Method == model.TaskMethodUploadAsset
}

func (t *Task) IsCmd() bool {
	return t.res.Task.Method == model.TaskMethodCmd
}

func (t *Task) Cmd() string {
	if t.IsCmd() {
		return t.res.Task.Cmd
	}
	return ""
}

func (u *UseCase) SubscribeTask(ctx context.Context) iter.Seq[Task] {
	return func(yield func(Task) bool) {
		times := 0

		for {
			t, err := u.adap.GetTask(ctx)
			if err != nil {
				u.AppDebug(ctx, "err: %s", err.Error())
				yield(Task{Err: ErrConnectOnGetTask})
				return
			}

			if t.Task.Status == model.TaskStatusProceeding {
				times = 0
				if !yield(Task{res: t}) {
					return
				}
			} else {
				times++
			}

			if times > 600 {
				yield(Task{Err: ErrConnectOnGetTask})
				return
			}
			time.Sleep(1 * time.Second)
		}
	}
}
