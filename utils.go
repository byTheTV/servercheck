package servercheck


import (
	"fmt"
	"net/url"
)

func ParseURL(urlStr string) (string, string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", "", err
	}
	host, port := u.Hostname(), u.Port()
	if port == "" {
		if u.Scheme == "http" {
			port = "80"
		} else if u.Scheme == "https" {
			port = "443"
		} else {
			return "", "", fmt.Errorf("unknown scheme: %s", u.Scheme)
		}
	}
	return host, port, nil
}