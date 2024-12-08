package mutation

import (
	"context"

	"github.com/enuesaa/taskhop/commander/gql/model"
)

func (r *MutationResolver) Log(ctx context.Context, input model.LogInput) (bool, error) {
	r.Logi.Info(context.Background(), input.Output)

	return true, nil
}
