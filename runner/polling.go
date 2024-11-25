package runner

import (
	"fmt"
	"net"
	"time"
)

func Polling() {
	for range 5 {
		address := net.JoinHostPort("localhost", "3000")
		conn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			fmt.Printf("接続失敗: %s\n", err)
		} else {
			fmt.Printf("接続成功: %s\n", address)
			conn.Close()
			break
		}
		time.Sleep(5 * time.Second)
	}
}
