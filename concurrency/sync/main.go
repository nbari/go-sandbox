package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/nbari/base58"
)

func main() {
	tasks := make(chan uint64, 1<<16)

	// spawn four worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for num := range tasks {
				x := base58.Encode(num)
				fmt.Printf("%d = %s\n", num, x)
				time.Sleep(time.Millisecond)
			}
			wg.Done()
		}()
	}

	// generate some tasks
	for i, val := uint64(0), uint64(1<<16); i < val; i++ {
		tasks <- i
	}
	close(tasks)

	// wait for the workers to finish
	wg.Wait()
}
