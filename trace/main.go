package main

import (
	"os"
	"time"
)

func main() {
	forever := make(chan struct{})
	go func() {
		for {
			os.Stderr.WriteString("message: foo bar")
			time.Sleep(3 * time.Second)
		}
	}()
	<-forever
}
