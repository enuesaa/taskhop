package mutation

import (
	"context"
)

func (r *MutationResolver) Register(ctx context.Context) (bool, error) {
	r.Lib.Proc.Register()
	r.Lib.Log.Info(context.Background(), "register!")

	return true, nil
}
