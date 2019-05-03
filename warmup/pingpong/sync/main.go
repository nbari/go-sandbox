package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)

	timer := time.After(5 * time.Second)

	pipo := func() {
		for {
			select {
			case <-ch:
				fmt.Println("ping")
				ch <- struct{}{}
				time.Sleep(time.Second)
				fmt.Println("pong\n")
			case <-timer:
				wg.Done()
			}
		}
	}

	go pipo()
	go pipo()

	ch <- struct{}{}

	wg.Wait()
}
