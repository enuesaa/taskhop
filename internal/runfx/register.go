package runfx

import "errors"

var ErrRegisterNotAvailable = errors.New("register not available")

func (i *Impl) Register() error {
	if i.Status != StatusWaiting {
		return ErrRegisterNotAvailable
	}
	i.Status = StatusProceeding

	return nil
}
