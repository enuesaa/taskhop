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
	IsCmd      bool
	Cmd        string
}

func (c *Connector) SubscribeTask(ctx context.Context) <-chan Task {
	ch := make(chan Task)

	go func() {
		defer close(ch)
		times := 0

		for {
			t, err := c.gql.GetTask(ctx)
			if err != nil {
				ch <- Task{
					Err: err,
				}
				break
			}

			if t.Task.Status == model.TaskStatusProceeding {
				times = 0
				if t.Task.Method == model.TaskMethodDownloadAsset {
					ch <- Task{
						IsDownload: true,
					}					
				} else {
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
