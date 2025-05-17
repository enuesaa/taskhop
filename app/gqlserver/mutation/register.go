package mutation

import (
	"context"

	"github.com/enuesaa/taskhop/lib/taskfx"
)

func (r *MutationResolver) Register(ctx context.Context) (bool, error) {
	id, err := r.Lib.Task.Register()
	if err != nil {
		return false, err
	}
	r.Lib.Log.AppInfo(context.Background(), "started: %s", id)

	task := r.Lib.Task.Get()
	if task.Status == taskfx.StatusPrompt || task.Status == taskfx.StatusDownloadAssets {
		go r.Lib.Task.Prompt()
	}

	return true, nil
}
