package connector

import (
	"context"
	"time"

	"github.com/enuesaa/taskhop/app/gql/model"
)

type Task struct {
	Err error
	Status model.TaskStatus
	Cmd string
}

func (c *Connector) SubscribeTask(ctx context.Context) <-chan Task {
	ch := make(chan Task)

	go func(){
		defer close(ch)

		for {
			taskres, err := c.Gql.GetTask(ctx)

			if err != nil {
				ch <- Task{
					Err: err,
				}
				break
			}
			ch <- Task{
				Status: taskres.Task.Status,
				Cmd: taskres.Task.Cmd,
			}
			time.Sleep(1 * time.Second)
		}
	}()

	return ch
}
