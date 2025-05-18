package gqlclient

import (
	"context"
	"fmt"
	"time"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
)

type Task struct {
	Err        error
	IsDownload bool
	IsUpload   bool
	IsCmd      bool
	Cmd        string
}

func (c *Connector) SubscribeTask(ctx context.Context) <-chan Task {
	ch := make(chan Task)

	go func() {
		defer close(ch)
		times := 0

		for {
			t, err := c.adap.GetTask()
			if err != nil {
				ch <- Task{
					Err: err,
				}
				break
			}

			if t.Task.Status == model.TaskStatusProceeding {
				times = 0
				switch t.Task.Method {
				case model.TaskMethodDownloadAsset:
					ch <- Task{
						IsDownload: true,
					}
				case model.TaskMethodUploadAsset:
					ch <- Task{
						IsUpload: true,
					}
				default:
					ch <- Task{
						IsCmd: true,
						Cmd:   t.Task.Cmd,
					}
				}
			} else {
				times++
			}

			if times > 600 {
				ch <- Task{
					Err: fmt.Errorf("limit"),
				}
				break
			}
			time.Sleep(1 * time.Second)
		}
	}()

	return ch
}
