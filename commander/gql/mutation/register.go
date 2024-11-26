package mutation

import (
	"context"
)

func (r *MutationResolver) Register(ctx context.Context) (bool, error) {
	r.Runi.Register()

	return true, nil
}
