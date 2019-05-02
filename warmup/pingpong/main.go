package main

import (
	"time"
)

func main() {
	ch := make(chan int)
	go player(ch)
	go player(ch)

	ch <- 1
	time.Sleep(time.Minute)
}

func player(ch chan int) {
	for {
		<-ch
		println("ping")
		time.Sleep(time.Second)
		ch <- 1
		println("pong\n")
	}
}
