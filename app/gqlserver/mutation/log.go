package mutation

import (
	"context"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
)

func (r *MutationResolver) Log(ctx context.Context, input model.LogInput) (bool, error) {
	r.li.Log.Faint(ctx, input.Output)

	return true, nil
}
