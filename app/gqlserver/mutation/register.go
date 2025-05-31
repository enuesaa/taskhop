package mutation

import (
	"context"

	"github.com/oklog/ulid/v2"
)

func (r *MutationResolver) Register(ctx context.Context) (string, error) {
	if err := r.li.Task.StartPrompt(); err != nil {
		return "", err
	}
	r.li.Log.Info(ctx, "started")
	id := ulid.Make().String()

	return id, nil
}
