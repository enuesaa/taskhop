package mutation

import "context"

func (r *MutationResolver) Register(ctx context.Context) (bool, error) {
	if err := r.li.Task.StartPrompt(); err != nil {
		return false, err
	}
	r.li.Log.AppInfo(ctx, "started")

	return true, nil
}
