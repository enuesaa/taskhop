package procfx

import "errors"

type Completed bool

var ErrCompletedNotAvailable = errors.New("unregister not available")

func (i *ProcSrv) NotifyCompleted() error {
	if i.status != StatusProceeding {
		return ErrCompletedNotAvailable
	}
	i.status = StatusCompleted
	i.completedCh <- Completed(true)

	return nil
}

func (i *ProcSrv) SubscribeCompleted() <-chan Completed {
	return i.completedCh
}
