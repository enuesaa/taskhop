package query

import (
	"context"
	"time"

	"github.com/enuesaa/taskhop/internal/routegql/schema"
)

func (r *QueryResolver) GetHealth(ctx context.Context) (*schema.Health, error) {
	data := schema.Health{
		Ok: true,
		Code: "aa",
		Time: time.Now().Local().Format(time.RFC3339),
	}

	return &data, nil
}
