package taskfx

import (
	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib/taskfx/repository"
)

func New(cl cli.ICli, repo repository.IRepository) ITaskSrv {
	return &TaskSrv{
		cli:    cl,
		repo:   repo,
	}
}

type ITaskSrv interface {
	Read() (Task, error)
	Validate(task Task) error
}

type TaskSrv struct {
	cli    cli.ICli
	repo   repository.IRepository
}
