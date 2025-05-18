package mutation

import "context"

func (r *MutationResolver) Completed(ctx context.Context) (bool, error) {
	if err := r.li.Task.MakeCompleted(); err != nil {
		return false, err
	}
	return true, nil
}
