package mutation

import (
	"context"

	"github.com/enuesaa/taskhop/app/gql/model"
)

func (r *MutationResolver) Log(ctx context.Context, input model.LogInput) (bool, error) {
	if input.Type == model.LogTypeCommand {
		return true, nil
	}
	r.Lib.Log.Faint(context.Background(), input.Output)

	return true, nil
}
