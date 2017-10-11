package main

import (
	"fmt"
	"time"
)

func ping(ch chan string) {
	for {
		fmt.Println(<-ch)
		ch <- "ping"
		time.Sleep(time.Second)
	}
}
func pong(ch chan string) {
	for {
		fmt.Println(<-ch)
		ch <- "pong"
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan string)
	go ping(ch)
	ch <- "toss"
	pong(ch)
	//time.Sleep(3 * time.Second)
}
