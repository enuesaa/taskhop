package connector

import (
	"errors"
	"time"
)

var ErrConnection = errors.New("failed to connect")

func (c *Client) DialPolling() error {
	for range 120 {
		if c.Dial() {
			return nil
		}
		time.Sleep(5 * time.Second)
	}
	return ErrConnection
}
