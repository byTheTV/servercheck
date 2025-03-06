package main

import (
	"net"
	"time"
)

func CheckUDP(host, port string) bool {
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
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	return err == nil
}
