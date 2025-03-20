package mutation

import (
	"context"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
)

func (r *MutationResolver) Log(ctx context.Context, input model.LogInput) (bool, error) {
	if input.Type == model.LogTypeCommand {
		r.Lib.Log.Faint(context.Background(), "")
		return true, nil
	}
	r.Lib.Log.Faint(context.Background(), input.Output)

	return true, nil
}
