package adapter

import (
	"context"
	"errors"
	"net"
	"time"
)

var ErrConnect = errors.New("failed to connect")

func (c *Adapter) Connect(ctx context.Context) error {
	if err := c.dialPolling(); err != nil {
		return err
	}
	if _, err := c.gql.GetHealth(ctx); err != nil {
		return err
	}
	if _, err := c.gql.Register(ctx); err != nil {
		return err
	}
	return nil
}

func (c *Adapter) dialPolling() error {
	for range 120 {
		if c.dial() {
			return nil
		}
		time.Sleep(5 * time.Second)
	}
	return ErrConnect
}

func (c *Adapter) dial() bool {
	conn, err := net.DialTimeout("tcp", c.address, 5*time.Second)
	if err != nil {
		return false
	}
	conn.Close()

	return true
}
