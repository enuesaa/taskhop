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

	if task := r.Lib.Task.Get(); task.Status == taskfx.StatusPrompt {
		go r.Lib.Task.Prompt()
	}

	return true, nil
}
