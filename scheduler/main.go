package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc(wg *sync.WaitGroup) {
	fmt.Printf("now: %+s\n", time.Now())
	time.Sleep(3 * time.Second)
	wg.Done()
}

func main() {
	quit := make(chan bool)

	t := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-t.C:
				var wg sync.WaitGroup
				for i := 0; i <= 4; i++ {
					wg.Add(1)
					go myFunc(&wg)
				}
				wg.Wait()
				fmt.Printf("--- done ---\n\n")
			case <-quit:
				return
			}
		}
	}()

	<-time.After(time.Minute)
	close(quit)
}
