package taskfx

import (
	"errors"

	"github.com/oklog/ulid/v2"
)

var ErrRegisterNotAvailable = errors.New("register not available")

func (i *TaskSrv) Register() (string, error) {
	if i.current.Status != StatusRegistration {
		return "", ErrRegisterNotAvailable
	}
	i.current.Status = StatusProceeding

	id := ulid.Make().String()

	return id, nil
}
