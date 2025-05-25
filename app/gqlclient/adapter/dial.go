package adapter

import (
	"net"
	"time"
)

func (c *Adapter) Dial() bool {
	conn, err := net.DialTimeout("tcp", c.address, 5*time.Second)
	if err != nil {
		return false
	}
	conn.Close() //nolint:errcheck

	return true
}
