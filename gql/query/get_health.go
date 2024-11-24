package query

import (
	"context"
	"time"

	"github.com/enuesaa/taskhop/gql/model"
)

func (r *QueryResolver) GetHealth(ctx context.Context) (*model.Health, error) {
	// readme, err := usecase.GetReadme(r.Repos)
	// if err != nil {
	// 	res := model.Health{
	// 		Ok:   false,
	// 		Code: "",
	// 		Time: time.Now().Local().Format(time.RFC3339),
	// 	}
	// 	return &res, nil
	// }

	res := model.Health{
		Ok:   true,
		Code: "",
		Time: time.Now().Local().Format(time.RFC3339),
	}
	return &res, nil
}
