package checkers

import (
	"net"
	"time"
)

type tcpChecker struct{}

func (c tcpChecker) Check(host, port string) bool {
	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func init() {
	RegisterChecker("tcp", tcpChecker{})
}