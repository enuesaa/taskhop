package taskfx

import "github.com/enuesaa/taskhop/internal/cli"

func New(config cli.Config) I {
	return &Impl{
		config: config,
	}
}

type I interface {
	Read() (Task, error)
	Validate(task Task) error
}

type Impl struct{
	config cli.Config
}
