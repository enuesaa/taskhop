package mutation

import "context"

func (r *MutationResolver) Register(ctx context.Context) (bool, error) {
	id, err := r.Lib.Proc.Register()
	if err != nil {
		return false, err
	}
	r.Lib.Log.Info(context.Background(), "runner registered %s", id)

	return true, nil
}
