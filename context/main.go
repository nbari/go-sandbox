package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func doWork(element int, wg *sync.WaitGroup) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()

	ch := make(chan struct{})

	go func(ch chan struct{}) {
		time.Sleep(time.Second)
		fmt.Printf("element = %+v\n", element)
		ch <- struct{}{}
	}(ch)

	select {
	case <-ch:
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	elements := []int{1, 2, 3}

	for _, element := range elements {
		wg.Add(1)
		go doWork(element, &wg)
	}
	wg.Wait()
}
