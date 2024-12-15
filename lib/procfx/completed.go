package procfx

import "errors"

var ErrCompletedNotAvailable = errors.New("unregister not available")

func (i *ProcSrv) Completed() error {
	if i.status != StatusProceeding {
		return ErrCompletedNotAvailable
	}
	i.status = StatusCompleted
	i.ch <- StatusCompleted

	return nil
}
