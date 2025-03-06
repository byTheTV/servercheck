package servercheck

import (
	"github.com/byTheTV/servercheck/checkers"
)

func CheckServer(urlStr string, protocols []string) (map[string]bool, error) {
	host, port, err := ParseURL(urlStr)
	if err != nil {
		return nil, err
	}
	return CheckServerWithHostPort(host, port, protocols), nil
}

func CheckServerWithHostPort(host, port string, protocols []string) map[string]bool {
	results := make(map[string]bool)
	for _, p := range protocols {
		if checker, ok := checkers.Checkers[p]; ok {
			results[p] = checker.Check(host, port)
		} else {
			results[p] = false
		}
	}
	return results
}
