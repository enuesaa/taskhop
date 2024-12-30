package procfx

import "github.com/enuesaa/taskhop/conf"

func New(config *conf.Config) IProcSrv {
	return &ProcSrv{
		config:      config,
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
	config      *conf.Config
	status      Status
	completedCh chan Completed
}
