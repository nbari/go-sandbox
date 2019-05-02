package main

import (
	"sync"
	"time"
)

func main() {
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	timer := time.After(5 * time.Second)

	// ping
	go func() {
		for {
			ch <- "ping"
		}
	}()

	// pong
	go func() {
		for {
			ch <- "pong"
		}
	}()

	go func() {
		for {
			select {
			case s := <-ch:
				println(s)
				time.Sleep(time.Second)
			case <-timer:
				wg.Done()
			}
		}
	}()

	wg.Wait()
}
