package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
				fmt.Println("--- done ---\n\n")
			case <-quit:
				return
			}
		}
	}()

	block := make(chan os.Signal)
	signal.Notify(block, syscall.SIGUSR1, syscall.SIGUSR2)
	<-block
	signal.Stop(block)
	log.Printf("%q signal received.", block)

}
