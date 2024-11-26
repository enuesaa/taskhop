package query

import (
	"context"

	"github.com/enuesaa/taskhop/commander/gql/model"
)

func (r *QueryResolver) Task(ctx context.Context) (*model.Task, error) {
	f, err := r.Taski.Read("testdata/cmds.yml")
	if err != nil {
		return nil, err
	}

	res := model.Task{
		Cmds: f.Cmds,
	}
	return &res, nil
}
