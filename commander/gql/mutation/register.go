package mutation

import (
	"context"
)

func (r *MutationResolver) Register(ctx context.Context) (bool, error) {
	r.Runi.Register()
	r.Logi.Info(context.Background(), "register!")

	return true, nil
}
