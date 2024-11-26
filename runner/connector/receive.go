package connector

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/enuesaa/taskhop/commander/gql/model"
	"github.com/enuesaa/taskhop/runner/client"
)

var ErrTaskNotAvailable = errors.New("tasl not available")

func (c *Connector) Receive(address string) (client.GetTask_Task, error) {
	url := fmt.Sprintf("http://%s/graphql", address)
	cli := client.NewClient(http.DefaultClient, url, nil)

	data, err := cli.GetTask(context.Background())
	if err != nil {
		return client.GetTask_Task{}, err
	}
	task := data.Task

	if task.Status != model.TaskStatusWaiting {
		return client.GetTask_Task{}, ErrTaskNotAvailable
	}
	return task, nil
}
