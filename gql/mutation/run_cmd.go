package mutation

import (
	"context"

	"github.com/enuesaa/taskhop/gql/model"
)

func (r *MutationResolver) RunCmd(ctx context.Context, input model.RunCmdInput) (bool, error) {
	return true, nil
}
