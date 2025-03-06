package checkers

import (
	"time"

	"github.com/go-ping/ping"
)

type icmpChecker struct{}

func (c icmpChecker) Check(host, port string) bool {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return false
	}

	pinger.Count = 3
	pinger.Timeout = 5 * time.Second
	err = pinger.Run()
	if err != nil {
		return false
	}

	stats := pinger.Statistics()
	return stats.PacketsRecv > 0
}

func init() {
	RegisterChecker("icmp", icmpChecker{})
}