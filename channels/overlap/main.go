package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type scheduler struct {
	t    <-chan time.Time
	quit chan struct{}
	f    func()
}

func myFunc() func() {
	return func() {
		start := time.Now()
		fmt.Printf("Start: %q ", start.Format(time.RFC3339))
		time.Sleep(time.Second * 4)
		fmt.Printf("Elapsed: %q - Gorutines: %d\n", time.Since(start), runtime.NumGoroutine())
	}
}

func main() {

	scheduler := scheduler{
		t:    time.NewTicker(time.Second * 1).C,
		quit: make(chan struct{}),
		f:    myFunc(),
	}

	_ = "breakpoint"

	go func() {
		for {
			select {
			case <-scheduler.t:
				scheduler.f()
				fmt.Printf("Gorutines: %d\n", runtime.NumGoroutine())
			case <-scheduler.quit:
				return
			}
		}
	}()

	// exit on signal
	block := make(chan os.Signal, 1)
	signal.Notify(block, os.Interrupt, os.Kill, syscall.SIGTERM)
	signalType := <-block
	fmt.Printf("%q signal received.", signalType)
	signal.Stop(block)
	os.Exit(0)
}
