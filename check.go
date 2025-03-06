
package servercheck

import (
	"github.com/byTheTV/servercheck/servercheck/checkers"
)

func CheckServer(host, port string, protocols []string) map[string]bool {
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