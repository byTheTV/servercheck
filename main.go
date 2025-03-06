package main

import (
	"fmt"
)

func main() {
	host, port, err := ParseURL("https://jsonplaceholder.typicode.com:443")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	results := CheckServer(host, port, []string{"http", "https", "tcp", "icmp", "udp"})
	alive := []string{}
	for p, ok := range results {
		if ok {
			alive = append(alive, p)
		}
	}

	if len(alive) > 0 {
		fmt.Printf("Доступен по: %v\n", alive)
	} else {
		fmt.Println("Сервер недоступен")
	}
}
