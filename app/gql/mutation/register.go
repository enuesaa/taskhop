package mutation

import "context"

func (r *MutationResolver) Register(ctx context.Context) (bool, error) {
	id, err := r.Lib.Task.Register()
	if err != nil {
		return false, err
	}
	r.Lib.Log.AppInfo(context.Background(), "started: %s", id)

	go r.Lib.Task.Prompt()

	return true, nil
}
