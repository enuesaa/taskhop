package query

import (
	"context"
	"time"

	"github.com/enuesaa/taskhop/internal/routegql/schema"
	"github.com/enuesaa/taskhop/internal/usecase"
)

func (r *QueryResolver) GetHealth(ctx context.Context) (*schema.Health, error) {
	readme, err := usecase.GetReadme(r.Repos)
	if err != nil {
		res := schema.Health{
			Ok:   false,
			Code: "",
			Time: time.Now().Local().Format(time.RFC3339),
		}
		return &res, nil
	}

	res := schema.Health{
		Ok:   true,
		Code: readme,
		Time: time.Now().Local().Format(time.RFC3339),
	}
	return &res, nil
}
