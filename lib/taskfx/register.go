package taskfx

import "errors"

var ErrRegisterNotAvailable = errors.New("register not available")

func (i *TaskSrv) Register() error {
	if i.status != StatusWaiting {
		return ErrRegisterNotAvailable
	}
	i.status = StatusProceeding

	return nil
}
