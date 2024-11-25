package mutation

import (
	"context"

	"github.com/enuesaa/taskhop/internal/runnermg"
)

func (r *MutationResolver) Register(ctx context.Context) (bool, error) {
	// TODO
	runnermg.New().Register()

	return true, nil
}
