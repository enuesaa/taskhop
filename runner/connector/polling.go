package connector

import (
	"fmt"
	"net"
	"time"
)

func (c *Connector) polling(commanderAddress string) error {
	for range 120 {
		conn, err := net.DialTimeout("tcp", commanderAddress, 5*time.Second)
		if err != nil {
			fmt.Printf("FAILED: %s\n", commanderAddress)
		} else {
			fmt.Printf("FOUND: %s\n", commanderAddress)
			conn.Close()
			return nil
		}
		time.Sleep(5 * time.Second)
	}

	return fmt.Errorf("failed to connect: %s", commanderAddress)
}
