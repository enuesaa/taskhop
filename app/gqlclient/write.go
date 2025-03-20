package gqlclient

import (
	"context"

	"github.com/enuesaa/taskhop/app/gqlserver/model"
)

// send logs
func (c *Connector) Write(b []byte) (int, error) {
	data := string(b)

	input := model.LogInput{
		Output: data,
	}
	_, err := c.gql.Log(context.Background(), input)
	if err != nil {
		return 0, err
	}

	return len(b), nil
}
