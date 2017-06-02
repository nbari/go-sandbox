package main

import (
	"io/ioutil"
	"time"
)

func main() {
	forever := make(chan struct{})
	go func() {
		for {
			ioutil.WriteFile("/dev/null", []byte("message: foo bar"), 644)
			time.Sleep(3 * time.Second)
		}
	}()
	<-forever
}
