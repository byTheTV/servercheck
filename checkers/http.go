package checkers

import (
	"net/http"
	"time"
)

type httpChecker struct {
	scheme string
}

func (c httpChecker) Check(host, port string) bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	url := c.scheme + "://" + host + ":" + port
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode >= 200 && resp.StatusCode < 400
}

func init() {
	RegisterChecker("http", httpChecker{scheme: "http"})
	RegisterChecker("https", httpChecker{scheme: "https"})
}
