package runfx

import "errors"

var ErrCompletedNotAvailable = errors.New("unregister not available")

func (i *Impl) Completed() error {
	if i.Status != StatusProceeding {
		return ErrCompletedNotAvailable
	}
	i.Status = StatusCompleted
	i.ch <- StatusCompleted

	return nil
}
