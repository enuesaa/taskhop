package connect

import (
	"fmt"
	"net"
	"time"
)

func polling(address string) error {
	for range 120 {
		conn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			fmt.Printf("FAILED: %s\n", address)
		} else {
			fmt.Printf("FOUND: %s\n", address)
			conn.Close()
			return nil
		}
		time.Sleep(5 * time.Second)
	}

	return fmt.Errorf("failed to connect: %s", address)
}