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

	if i.config.TransferFlag && !i.assetsDownloaded {
		i.current.Status = StatusDownloadAssets
	} else {
		i.current.Status = StatusPrompt
	}
	id := ulid.Make().String()

	return id, nil
}
