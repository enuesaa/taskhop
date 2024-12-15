package procfx

import (
	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib/taskfx/repository"
)

func New(cl cli.ICli, repo repository.IRepository) IProcSrv {
	return &ProcSrv{
		cli:    cl,
		status: StatusWaiting,
		ch:     make(chan Status, 1),
		repo:   repo,
	}
}

type IProcSrv interface {
	GetStatus() Status
	Register() error
	Completed() error
	Subscribe() <-chan Status
}

type ProcSrv struct {
	cli    cli.ICli
	status Status
	ch     chan Status
	repo   repository.IRepository
}
