package mutation

import (
	"context"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
)

func (r *MutationResolver) Log(ctx context.Context, input model.LogInput) (bool, error) {
	r.Lib.Log.Faint(context.Background(), input.Output)

	return true, nil
}
