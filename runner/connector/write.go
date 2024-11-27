package connector

import (
	"context"

	"github.com/enuesaa/taskhop/commander/gql/model"
)

// send logs
func (c *Client) Write(b []byte) (int, error) {
	data := string(b)

	input := model.LogInput{
		Output: data,
	}
	_, err := c.Log(context.Background(), input)
	if err != nil {
		return 0, err
	}

	return len(b), nil
}
