package ping

import (
	"fmt"
	"net"
	"time"
)

func CheckTCPPort(address string, port int, timeout time.Duration) error {
	target := fmt.Sprintf("%s:%d", address, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
