package connector

import (
	"net"
	"time"
)

func (c *Client) Dial() bool {
	conn, err := net.DialTimeout("tcp", c.Address(), 5 * time.Second)
	if err != nil {
		return false
	}
	conn.Close()

	return true
}
