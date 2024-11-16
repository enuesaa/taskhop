package query

import (
	"context"

	"github.com/enuesaa/taskhop/internal/routegql/schema"
)

func (r *QueryResolver) GetHealth(ctx context.Context) (*schema.Health, error) {
	data := schema.Health{}

	return &data, nil
}
