package gqlclient

import (
	"context"
	"errors"
	"net"
	"time"
)

var ErrConnect = errors.New("failed to connect")

func (c *Connector) Connect(ctx context.Context) error {
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

func (c *Connector) dialPolling() error {
	for range 120 {
		if c.dial() {
			return nil
		}
		time.Sleep(5 * time.Second)
	}
	return ErrConnect
}

func (c *Connector) dial() bool {
	conn, err := net.DialTimeout("tcp", c.config.Address, 5*time.Second)
	if err != nil {
		return false
	}
	conn.Close()

	return true
}
