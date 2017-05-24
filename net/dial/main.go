package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(254)
	timeout := 2 * time.Second
	for i := 1; i < 255; i++ {
		go func(i int) {
			defer wg.Done()
			host := fmt.Sprintf("192.168.1.%d:80", i)
			if conn, err := net.DialTimeout("tcp", host, timeout); err == nil {
				fmt.Printf("online = %s\n", host)
				conn.Close()
			}
		}(i)
	}
	wg.Wait()
}
