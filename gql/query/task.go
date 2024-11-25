package query

import (
	"context"

	"github.com/enuesaa/taskhop/gql/model"
	"github.com/enuesaa/taskhop/internal/cmdsfile"
)

func (r *QueryResolver) Task(ctx context.Context) (*model.Task, error) {
	// TODO
	f, err := cmdsfile.New().Read("testdata/cmds.yml")
	if err != nil {
		return nil, err
	}

	res := model.Task{
		Cmds: f.Cmds,
	}
	return &res, nil
}
