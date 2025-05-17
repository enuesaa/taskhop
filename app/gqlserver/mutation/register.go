package mutation

import "context"

func (r *MutationResolver) Register(ctx context.Context) (bool, error) {
	if err := r.Lib.Task.StartPrompt(); err != nil {
		return false, err
	}
	r.Lib.Log.AppInfo(ctx, "started")

	return true, nil
}
