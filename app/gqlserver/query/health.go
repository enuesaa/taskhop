package query

import (
	"context"
	"time"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
)

func (r *QueryResolver) Health(ctx context.Context) (*model.Health, error) {
	res := model.Health{
		Ok:   true,
		Code: "",
		Time: time.Now().Local().Format(time.RFC3339),
	}
	return &res, nil
}
