package procfx

import "errors"

var ErrRegisterNotAvailable = errors.New("register not available")

func (i *ProcSrv) Register() error {
	if i.status != StatusWaiting {
		return ErrRegisterNotAvailable
	}
	i.status = StatusProceeding

	return nil
}
