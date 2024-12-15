package taskfx

import (
	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib/taskfx/repository"
)

func New(cl cli.ICli, repo repository.IRepository) ITaskSrv {
	return &TaskSrv{
		cli:    cl,
		status: StatusWaiting,
		ch:     make(chan Status, 1),
		repo:   repo,
	}
}

type ITaskSrv interface {
	Read() (Task, error)
	Validate(task Task) error
	GetStatus() Status
	Register() error
	Completed() error
	Subscribe() <-chan Status
}

type TaskSrv struct {
	cli    cli.ICli
	status Status
	ch     chan Status
	repo   repository.IRepository
}
