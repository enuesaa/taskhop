package mutation

import (
	"context"

	"github.com/enuesaa/taskhop/app/commander/gql/model"
)

func (r *MutationResolver) Log(ctx context.Context, input model.LogInput) (bool, error) {
	r.Lib.Log.Info(context.Background(), input.Output)

	return true, nil
}
