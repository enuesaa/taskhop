package mutation

import (
	"context"

	"github.com/enuesaa/taskhop/commander/gql/model"
	"github.com/enuesaa/taskhop/internal/logging"
)

func (r *MutationResolver) Log(ctx context.Context, input model.LogInput) (bool, error) {
	// TODO
	logging.New().Info(context.Background(), input.Output)

	return true, nil
}
