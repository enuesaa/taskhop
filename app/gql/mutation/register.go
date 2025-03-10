package mutation

import "context"

func (r *MutationResolver) Register(ctx context.Context) (bool, error) {
	id, err := r.Lib.Proc.Register()
	if err != nil {
		return false, err
	}
	r.Lib.Log.AppInfo(context.Background(), "started: %s", id)

	return true, nil
}
