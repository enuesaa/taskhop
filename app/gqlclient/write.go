package gqlclient

import (
	"context"

	"github.com/enuesaa/taskhop/app/gql/model"
)

// send logs
func (c *Connector) Write(b []byte) (int, error) {
	data := string(b)

	input := model.LogInput{
		Type:   model.LogTypeCommandOutput,
		Output: data,
	}
	_, err := c.Gql.Log(context.Background(), input)
	if err != nil {
		return 0, err
	}

	return len(b), nil
}
