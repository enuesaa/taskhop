package taskfx

import (
	"github.com/enuesaa/taskhop/internal/cli"
	"github.com/enuesaa/taskhop/internal/taskfx/repository"
)

func New(config cli.Config, repo repository.I) I {
	return &Impl{
		config: config,
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
	config cli.Config
	status Status
	ch chan Status
	repo repository.I
}
