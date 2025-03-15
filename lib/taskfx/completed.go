package taskfx

import "errors"

type Completed bool

var ErrCompletedNotAvailable = errors.New("unregister not available")

func (i *TaskSrv) NotifyCompleted() error {
	if i.current.Status != StatusProceeding {
		return ErrCompletedNotAvailable
	}
	i.current.Status = StatusCompleted

	return nil
}
