package connector

import (
	"context"

	"github.com/enuesaa/taskhop/app/commander/gql/model"
)

// send logs
func (c *Connector) Write(b []byte) (int, error) {
	data := string(b)

	input := model.LogInput{
		Type:   model.LogTypeCommandOutput,
		Output: data,
	}
	_, err := c.Log(context.Background(), input)
	if err != nil {
		return 0, err
	}

	return len(b), nil
}
