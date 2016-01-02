package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type scheduler struct {
	t    <-chan time.Time
	quit chan struct{}
	f    func()
}

func main() {

	scheduler := scheduler{
		t:    time.NewTicker(time.Second * 1).C,
		quit: make(chan struct{}),
		f: func() {
			start := time.Now()
			fmt.Printf("Start: %q ", start.Format(time.RFC3339))
			time.Sleep(time.Second * 4)
			fmt.Printf("Elapsed: %q\n", time.Since(start))
		},
	}

	go func() {
		for {
			select {
			case <-scheduler.t:
				scheduler.f()
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
