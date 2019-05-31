package main

import "time"

func main() {
	ch := make(chan struct{})

	// run something blocking
	go func() {
		time.Sleep(5 * time.Second)
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		println("task done")
	// timeout after 2 seconds
	case <-time.After(2 * time.Second):
		println("timeout after 2 seconds")
	}
}
