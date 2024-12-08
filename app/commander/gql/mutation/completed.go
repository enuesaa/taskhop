package mutation

import "context"

func (r *MutationResolver) Completed(ctx context.Context) (bool, error) {
	if err := r.Taski.Completed(); err != nil {
		return false, err
	}

	return true, nil
}
