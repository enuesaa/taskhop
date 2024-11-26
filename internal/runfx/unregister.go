package runfx

import "errors"

var ErrUnRegisterNotAvailable = errors.New("unregister not available")

func (i *Impl) UnRegister() error {
	if i.Status != StatusProceeding {
		return ErrUnRegisterNotAvailable
	}
	i.Status = StatusCompleted

	return nil
}
