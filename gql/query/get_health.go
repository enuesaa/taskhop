package query

import (
	"context"
	"time"

	"github.com/enuesaa/taskhop/gql/model"
	"github.com/enuesaa/taskhop/internal/task"
)

func (r *QueryResolver) GetHealth(ctx context.Context) (*model.Health, error) {
	readme, err := task.GetReadme(r.Fs)
	if err != nil {
		res := model.Health{
			Ok:   false,
			Code: "",
			Time: time.Now().Local().Format(time.RFC3339),
		}
		return &res, nil
	}

	res := model.Health{
		Ok:   true,
		Code: readme,
		Time: time.Now().Local().Format(time.RFC3339),
	}
	return &res, nil
}
