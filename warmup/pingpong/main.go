package main

import (
	"time"
)

func ping(ch chan int) {
	for {
		ball := <-ch
		ball++
		print(ball)
		ch <- ball
	}
}
func pong(ch chan int) {
	for {
		ball := <-ch
		ball--
		print(ball)
		ch <- ball
	}
}

func main() {
	ch := make(chan int)
	go ping(ch)
	go pong(ch)
	ch <- 1
	time.Sleep(time.Second)
}
