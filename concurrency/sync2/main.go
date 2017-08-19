package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/nbari/base58"
)

const maxConcurrency = 4

var throttle = make(chan struct{}, maxConcurrency)

func main() {
	var wg sync.WaitGroup
	for i, val := uint64(0), uint64(1<<16); i < val; i++ {
		throttle <- struct{}{}
		wg.Add(1)
		go f(i, &wg, throttle)
	}
	wg.Wait()
}

func f(i uint64, wg *sync.WaitGroup, throttle chan struct{}) {
	defer wg.Done()
	x := base58.Encode(i)
	fmt.Printf("%d = %s\n", i, x)
	time.Sleep(time.Millisecond)
	<-throttle
}
