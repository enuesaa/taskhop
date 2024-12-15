package procfx

import (
	"errors"

	"github.com/oklog/ulid/v2"
)

var ErrRegisterNotAvailable = errors.New("register not available")

func (i *ProcSrv) Register() (string, error) {
	if i.status != StatusWaiting {
		return "", ErrRegisterNotAvailable
	}
	i.status = StatusProceeding

	id := ulid.Make().String()

	return id, nil
}
