package main

import "time"

func ping(ping <-chan int, pong chan<- int) {
	for {
		<-ping
		println("ping")
		time.Sleep(time.Second)
		pong <- 1
	}
}

func pong(ping chan<- int, pong <-chan int) {
	for {
		<-pong
		println("pong")
		time.Sleep(time.Second)
		ping <- 1
	}
}

func main() {
	pi := make(chan int)
	po := make(chan int)

	go ping(pi, po)
	go pong(pi, po)
	pi <- 1

	for {
		time.Sleep(time.Second)
	}
}
