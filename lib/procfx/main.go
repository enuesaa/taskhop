package procfx

import "github.com/enuesaa/taskhop/cli"

func New(cl cli.ICli) IProcSrv {
	return &ProcSrv{
		cli:         cl,
		status:      StatusWaiting,
		completedCh: make(chan Completed, 1),
	}
}

type IProcSrv interface {
	GetStatus() Status
	Register() (string, error)
	NotifyCompleted() error
	SubscribeCompleted() <-chan Completed
}

type ProcSrv struct {
	cli         cli.ICli
	status      Status
	completedCh chan Completed
}
