package taskfx

import "errors"

type Completed bool

var ErrCompletedNotAvailable = errors.New("unregister not available")

func (i *TaskSrv) NotifyCompleted() error {
	if i.status != StatusProceeding {
		return ErrCompletedNotAvailable
	}
	i.status = StatusCompleted

	return nil
}
