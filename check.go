package main

// CheckServer проверяет доступность сервера по указанным протоколам
func CheckServer(host, port string, protocols []string) map[string]bool {
	results := make(map[string]bool)
	for _, p := range protocols {
		results[p] = isAlive(p, host, port)
	}
	return results
}

func isAlive(protocol, host, port string) bool {
	switch protocol {
	case "http", "https":
		return CheckHTTP(protocol + "://" + host + ":" + port)
	case "tcp":
		return CheckTCP(host, port)
	case "udp":
		return CheckUDP(host, port)
	case "icmp":
		return CheckICMP(host)
	default:
		return false
	}
}