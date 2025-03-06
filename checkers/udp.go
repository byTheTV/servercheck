package checkers

import (
	"net"
	"time"
)

type udpChecker struct{}

func (c udpChecker) Check(host, port string) bool {
	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("udp", address, 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()

	_, err = conn.Write([]byte("ping"))
	if err != nil {
		return false
	}

	deadline := time.Now().Add(5 * time.Second)
	err = conn.SetReadDeadline(deadline)
	if err != nil {
		return false
	}

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		return false
	}

	return true
}

func init() {
	RegisterChecker("udp", udpChecker{})
}
