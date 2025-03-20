package gqlclient

import (
	"context"
	"fmt"
	"time"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
)

type Task struct {
	Err error
	IsDownload bool
	IsCmd bool
	Cmd string
}

func (c *Connector) SubscribeTask(ctx context.Context) <-chan Task {
	ch := make(chan Task)

	go func(){
		defer close(ch)
		times := 0

		for {
			taskres, err := c.Gql.GetTask(ctx)
			if err != nil {
				ch <- Task{
					Err: err,
				}
				break
			}

			switch taskres.Task.Status {
			case model.TaskStatusDownloadAssets:
				ch <- Task{
					IsDownload: true,
				}
				times = 0
			case model.TaskStatusProceeding:
				ch <- Task{
					IsCmd: true,
					Cmd: taskres.Task.Cmd,
				}
				times = 0
			default:
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
