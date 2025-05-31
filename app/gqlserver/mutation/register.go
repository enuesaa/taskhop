package mutation

import (
	"context"
	"math/rand"
	"fmt"
)

func (r *MutationResolver) Register(ctx context.Context) (string, error) {
	id := fmt.Sprintf("%05x", rand.Intn(0xFFFFF))
	if err := r.li.Task.StartPrompt(); err != nil {
		return "", err
	}
	r.li.Log.Info(ctx, "started")

	return id, nil
}
