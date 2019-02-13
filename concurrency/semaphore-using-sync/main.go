package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/nbari/base58"
)

func encode(i uint64) {
	x := base58.Encode(i)
	fmt.Printf("%d = %s\n", i, x)
	time.Sleep(time.Second)
}

func main() {
	var wg = sync.WaitGroup{}
	concurrency := 4
	sem := make(chan struct{}, concurrency)
	for i, val := uint64(0), uint64(1<<4); i <= val; i++ {
		sem <- struct{}{}
		wg.Add(1)
		go func(i uint64) {
			defer func() {
				<-sem
				wg.Done()
			}()
			encode(i)
		}(i)
	}

	wg.Wait()
}
