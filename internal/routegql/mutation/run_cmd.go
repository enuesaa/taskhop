package mutation

import (
	"context"

	"github.com/enuesaa/taskhop/internal/routegql/schema"
)

func (r *MutationResolver) RunCmd(ctx context.Context, input schema.RunCmdInput) (bool, error) {
	return true, nil
}
