package usecase

import (
	"context"

	"github.com/enuesaa/taskhop/commander/gql/model"
	"github.com/enuesaa/taskhop/runner/gqlclient"
)

func (c *UseCase) Run(task gqlclient.GetTask_Task, workdir string) error {
	w := LogWriter{
		Callback: func (data string) error {
			input := model.LogInput{
				Output: data,
			}
			_, err := c.client.Log(context.Background(), input)
			if err != nil {
				return err
			}
			return nil
		},
	}

	if err := c.Cmdi.MultiExec(&w, task.Cmds, workdir); err != nil {
		return err
	}

	return nil
}
