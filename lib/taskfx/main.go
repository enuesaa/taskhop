package taskfx

import (
	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib/taskfx/repository"
)

func New(cl cli.ICli, repo repository.I) I {
	return &Impl{
		cli: cl,
		status: StatusWaiting,
		ch: make(chan Status, 1),
		repo: repo,
	}
}

type I interface {
	Read() (Task, error)
	Validate(task Task) error
	GetStatus() Status
	Register() error
	Completed() error
	Subscribe() <-chan Status
}

type Impl struct{
	cli cli.ICli
	status Status
	ch chan Status
	repo repository.I
}
