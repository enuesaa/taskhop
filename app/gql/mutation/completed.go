package mutation

import "context"

func (r *MutationResolver) Completed(ctx context.Context) (bool, error) {
	if err := r.Lib.Task.NotifyCompleted(); err != nil {
		return false, err
	}

	go r.Lib.Task.Wait()

	return true, nil
}
