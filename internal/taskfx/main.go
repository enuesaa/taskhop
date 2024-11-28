package taskfx

import "github.com/enuesaa/taskhop/internal/cli"

func New(config cli.Config) I {
	return &Impl{
		config: config,
		status: StatusWaiting,
		ch: make(chan Status, 1),
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
}
