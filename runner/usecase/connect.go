package usecase

import (
	"context"
	"errors"
	"time"
)

var ErrConnection = errors.New("failed to connect")

func (c *UseCase) Connect() error {
	// polling
	for range 120 {
		if c.client.Dial() {
			break
		}
		time.Sleep(5 * time.Second)
	}

	data, err := c.client.GetHealth(context.Background())
	if err != nil {
		return ErrConnection
	}
	if !data.Health.Ok {
		return ErrConnection
	}
	return nil
}
