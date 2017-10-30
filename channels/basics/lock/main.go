package main

import "time"

func main() {
	run := make(chan string, 1)
	go func() {
		time.Sleep(time.Second)
		run <- "ok"
	}()
	<-run
}
