package taskfx

import (
	"errors"
	"time"

	"github.com/oklog/ulid/v2"
)

var ErrRegisterNotAvailable = errors.New("register not available")
var ErrRunnerUnHealthy = errors.New("runner unhealthy")

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

	go func() {
		for {
			time.Sleep(5 * time.Second)

			if time.Since(i.lastHealthy) > 5*time.Second {
				i.errch <- ErrRunnerUnHealthy
				break
			}
		}
	}()

	return id, nil
}
