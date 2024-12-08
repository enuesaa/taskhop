package connector

import (
	"net"
	"time"
)

func (c *Connector) Dial() bool {
	conn, err := net.DialTimeout("tcp", c.config.Address, 5 * time.Second)
	if err != nil {
		return false
	}
	conn.Close()

	return true
}
